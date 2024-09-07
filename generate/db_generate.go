package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/db"
	"github.com/kolo7/bench-tpl/varmanager"
	"github.com/pkg/errors"
)

type DBGenerator struct {
	cfg *config.Config

	varManager   *varmanager.VarManager
	schemaParser db.SchemaParser
}

func NewDBGenerator(cfg *config.Config) *DBGenerator {
	g := &DBGenerator{
		cfg: cfg,
	}

	conn := db.NewDB(cfg.DB.Dsn)
	g.schemaParser = db.NewSchemaParser(conn, cfg)
	g.varManager = varmanager.NewVarManager()
	return g
}

func (g *DBGenerator) Generate(ctx context.Context) error {
	tables, err := g.schemaParser.Parse()
	if err != nil {
		return err
	}
	g.varManager.SetGlobalVar("db", tables)

	for _, table := range tables {
		tableGenerate := NewGenerator(g.cfg, g.cfg.Tables[table.Name], g.varManager, table)
		text, err := tableGenerate.Generate(ctx)
		if err != nil {
			return err
		}
		err = g.outputFile(ctx, table.Name, text)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *DBGenerator) outputFile(ctx context.Context, tableName string, text string) error {
	outputFile := fmt.Sprintf("%s/%s.%s", g.cfg.Output.Dir, tableName, g.cfg.Output.Format)
	file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		// 创建目录
		dir := g.cfg.Output.Dir
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
