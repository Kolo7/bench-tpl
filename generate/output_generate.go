package generate

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/varmanager"
	"github.com/pkg/errors"
)

type TplGenerator struct {
	name       string
	cfg        *config.Config
	tplCfg     *config.TplConf
	varManager *varmanager.VarManager

	defaultEpoch int
	tplFileName  string

	tpl *template.Template
}

func NewGenerator(
	name string,
	cfg *config.Config,
	tplCfg *config.TplConf,
	varManager *varmanager.VarManager) *TplGenerator {
	g := &TplGenerator{
		name:         name,
		cfg:          cfg,
		tplCfg:       tplCfg,
		varManager:   varManager,
		defaultEpoch: tplCfg.Epoch,
		tplFileName:  fmt.Sprintf("%s/%s.%s", tplCfg.Dir, name, tplCfg.Format),
	}

	return g
}

func (g *TplGenerator) Generate(ctx context.Context) (string, error) {
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
		data := g.varManager.GetTables()
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
func (g *TplGenerator) loadTemplate() ([]byte, error) {
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
