package generate

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/Kolo7/bench-tpl/config"
	"github.com/Kolo7/bench-tpl/varmanager"
	"github.com/pkg/errors"
)

type EpochGenerator struct {
	name       string
	cfg        *config.Config
	epochCfg   *config.EpochConf
	varManager *varmanager.VarManager

	defaultEpoch int
	tplFileName  string

	tpl *template.Template
}

func NewEpochGenerator(
	name string,
	cfg *config.Config,
	epochCfg *config.EpochConf,
	varManager *varmanager.VarManager) Generator {
	g := &EpochGenerator{
		name:         name,
		cfg:          cfg,
		epochCfg:     epochCfg,
		varManager:   varManager,
		defaultEpoch: epochCfg.Epoch,
		tplFileName:  fmt.Sprintf("%s/%s.%s", epochCfg.Dir, name, epochCfg.Format),
	}

	return g
}

func (g *EpochGenerator) Generate(ctx context.Context) (string, error) {
	var (
		outText strings.Builder
	)

	content, err := g.loadTemplate()
	if err != nil {
		return "", err
	}
	tpl := template.New(g.name)

	funcs := g.varManager.GetFuncMap()
	tpl = tpl.Funcs(funcs)
	g.tpl, err = tpl.Parse(string(content))
	if err != nil {
		return "", errors.Wrapf(err, "解析模板文件 %s 失败", g.tplFileName)
	}

	for i := 0; i < g.defaultEpoch; i++ {
		data := g.varManager.GetTablesExampleVar()
		err = tpl.Execute(&outText, data)
		if err != nil {
			return "", errors.Wrapf(err, "生成 %s 模板失败", g.name)
		}
		if i < g.defaultEpoch-1 {
			outText.WriteString("\n")
		}

	}

	return outText.String(), nil
}

// 加载模板文件
func (g *EpochGenerator) loadTemplate() ([]byte, error) {
	file, err := os.OpenFile(g.tplFileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, errors.Wrapf(err, "打开模板文件 %s 失败", g.tplFileName)
	}
	defer file.Close()
	var content []byte
	// 用流的方式读取模板文件内容
	if content, err = io.ReadAll(file); err != nil {
		return nil, errors.Wrapf(err, "读取模板文件 %s 失败", g.tplFileName)
	}
	return content, nil
}
