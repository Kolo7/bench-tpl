package generate

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/db"
	"github.com/kolo7/bench-tpl/varmanager"
	"github.com/pkg/errors"
)

type TableGenerator struct {
	varManager *varmanager.VarManager

	table *db.Table
	cfg   config.TableConf

	defaultEpoch  int
	defaultOutput string
}

func NewGenerator(cfg config.TableConf, varManager *varmanager.VarManager, table *db.Table) *TableGenerator {
	g := &TableGenerator{
		varManager:    varManager,
		cfg:           cfg,
		table:         table,
		defaultEpoch:  100,
		defaultOutput: fmt.Sprintf("%s.json", table.Name),
	}

	return g
}

func (g *TableGenerator) Generate(ctx context.Context) error {
	tpl := template.New(g.table.Name)
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(g.defaultOutput, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	err = tpl.Funcs(g.varManager.Get()).Execute(file, g.table)
	if err != nil {
		return errors.Wrapf(err, "生成 %s 模板失败", g.table.Name)
	}
	return nil
}
