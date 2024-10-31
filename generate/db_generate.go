package generate

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/Kolo7/bench-tpl/config"
	"github.com/Kolo7/bench-tpl/db"
	"github.com/Kolo7/bench-tpl/input"
	"github.com/Kolo7/bench-tpl/varmanager"
	"github.com/pkg/errors"
)

type DBGenerator struct {
	cfg *config.Config

	varManager   *varmanager.VarManager
	schemaParser db.SchemaParser
	tplFileInput *input.TplFileInput
}

func NewDBGenerator(cfg *config.Config) Generator {
	g := &DBGenerator{
		cfg: cfg,
	}

	conn := db.NewDB(cfg.Dsn)
	g.schemaParser = db.NewSchemaParser(conn, cfg)
	g.varManager = varmanager.NewVarManager()
	g.tplFileInput = input.NewTplFileInput(cfg)
	return g
}

func (g *DBGenerator) Generate(ctx context.Context) (string, error) {

	for name, epoch := range g.cfg.EpochConf {
		tableGenerate := NewEpochGenerator(name, g.cfg, epoch, g.varManager)
		text, err := tableGenerate.Generate(ctx)
		if err != nil {
			return "", err
		}
		err = g.outputFile(ctx, name, text)
		if err != nil {
			return "", err
		}
	}
	tables, err := g.schemaParser.Parse()
	if err != nil {
		return "", err
	}
	// g.varManager.SetGlobalVar("db", tables)

	for _, table := range tables {
		g.varManager.SetTableExampleVar(table)
	}

	tpl := template.New("").Funcs(g.varManager.GetFuncMap())
	tpl, err = g.tplFileInput.LoadTemplate(tpl, "")
	if err != nil {
		return "", err
	}

	// 设置全局变量
	g.varManager.SetGlobalVar(g.cfg.FQDN)
	for name := range tables {
		nest := NewNestGenerator(g.cfg, name, g.varManager, tpl)
		if _, err := nest.Generate(ctx); err != nil {
			return "", err
		}
	}
	GoFmt(g.cfg.OutputDir)
	return "", nil
}

func (g *DBGenerator) outputFile(ctx context.Context, tplName string, text string) error {
	outputFile := fmt.Sprintf("%s/%s.%s", g.cfg.OutputDir, tplName, "json")
	file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		// 创建目录
		dir := g.cfg.OutputDir
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return errors.Wrapf(err, "创建目录 %s 失败", dir)
			}
		}
		file, err = os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return errors.Wrapf(err, "创建文件 %s 失败", outputFile)
		}
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return errors.Wrapf(err, "写入文件 %s 失败", outputFile)
	}
	return nil
}

// GoFmt exec gofmt for a code dir
func GoFmt(codeDir string) {
	args := []string{"-s", "-d", "-w", "-l", codeDir}
	cmd := exec.Command("gofmt", args...)
	cmd.Run()
}
