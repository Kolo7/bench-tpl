package generate

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/db"
	"github.com/kolo7/bench-tpl/varmanager"
	"github.com/pkg/errors"
)

type TableGenerator struct {
	cfg        *config.Config
	tableCfg   *config.TableConf
	varManager *varmanager.VarManager
	table      *db.Table

	defaultEpoch int
	tplFileName  string

	tpl *template.Template
}

func NewGenerator(
	cfg *config.Config,
	tableCfg *config.TableConf,
	varManager *varmanager.VarManager,
	table *db.Table) *TableGenerator {
	g := &TableGenerator{
		cfg:          cfg,
		tableCfg:     tableCfg,
		varManager:   varManager,
		table:        table,
		defaultEpoch: tableCfg.Epoch,
		tplFileName:  fmt.Sprintf("%s/%s.%s", cfg.Input.Dir, table.Name, cfg.Input.Format),
	}

	return g
}

func (g *TableGenerator) Generate(ctx context.Context) (string, error) {
	var (
		outText strings.Builder
	)

	content, err := g.loadTemplate()
	if err != nil {
		return "", err
	}
	if err = g.repeatTemplate(content); err != nil {
		return "", err
	}

	tpl := g.tpl.Funcs(template.FuncMap{})
	data := g.varManager.Get()
	err = tpl.Execute(&outText, data)
	if err != nil {
		return "", errors.Wrapf(err, "生成 %s 模板失败", g.table.Name)
	}
	return outText.String(), nil
}

// 读取模板文件,用epoch变量将模板重复n次
func (g *TableGenerator) repeatTemplate(content []byte) error {
	var (
		epoch = g.defaultEpoch
	)
	if g.tableCfg.Epoch == 0 {
		epoch = 1
	}

	// 用epoch变量将模板重复n行
	content = []byte(fmt.Sprintf("%s\n%s", string(content), strings.Repeat(string(content), epoch-1)))
	// 用新的模板内容创建模板对象
	tpl, err := template.New(g.table.Name).Parse(string(content))
	if err != nil {
		return errors.Wrapf(err, "解析模板文件 %s 失败", g.tplFileName)
	}
	g.tpl = tpl
	return nil
}

// 加载模板文件
func (g *TableGenerator) loadTemplate() ([]byte, error) {
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
