package generate

import (
	"context"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/db"
	"github.com/kolo7/bench-tpl/varmanager"
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
	g.varManager.SetVar("db", tables)

	for _, table := range tables {
		tableGenerate := NewGenerator(g.cfg.Tables[table.Name], g.varManager, table)
		err = tableGenerate.Generate(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
