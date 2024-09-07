package db

import (
	"github.com/kolo7/bench-tpl/config"
	"github.com/pkg/errors"
)

type SchemaParser interface {
	Parse() (map[string]*Table, error)
}

type defaultSchemaParser struct {
	cfg *config.Config
	DB  *DB
}

func NewSchemaParser(db *DB, cfg *config.Config) SchemaParser {
	return &defaultSchemaParser{
		cfg: cfg,
		DB:  db,
	}
}

func (p *defaultSchemaParser) Parse() (map[string]*Table, error) {
	tables := make(map[string]*Table)
	for tableName := range p.cfg.Tables {
		table, err := p.parseTable(tableName)
		if err != nil {
			return nil, errors.Wrapf(err, "解析表结构失败: %s", tableName)
		}
		tables[tableName] = table
	}
	return tables, nil
}

func (p *defaultSchemaParser) parseTable(tableName string) (*Table, error) {
	// 查询表结构
	v := make([]*Column, 0)
	err := p.DB.QueryRow(&v, "DESCRIBE "+tableName)
	if err != nil {
		return nil, errors.Wrapf(err, "查询表结构失败: %s", tableName)
	}
	rowResult := make([][]*Row, 0)
	rawDB, err := p.DB.SqlConn.RawDB()
	if err != nil {
		return nil, errors.Wrapf(err, "获取原始数据库连接失败: %s", tableName)
	}
	rows, err := rawDB.Query("SELECT * FROM " + tableName)
	if err != nil {
		return nil, errors.Wrapf(err, "查询表数据失败: %s", tableName)
	}
	defer rows.Close()
	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			return nil, errors.Wrapf(err, "获取列名失败: %s", tableName)
		}
		vals := make([]*Row, len(cols))
		for i := range vals {
			vals[i] = new(Row)
			vals[i].Name = cols[i]
			err = rows.Scan(vals[i])
			if err != nil {
				return nil, errors.Wrapf(err, "扫描行数据失败: %s", tableName)
			}
		}
		rowResult = append(rowResult, vals)
	}
	// 解析表结构
	return &Table{
		Name:    tableName,
		Columns: v,
		Rows:    rowResult,
	}, nil
}
