package db

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/Kolo7/bench-tpl/config"
	"github.com/Kolo7/bench-tpl/utils"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type SchemaParser interface {
	Parse() (map[string]*Table, error)
}

type defaultSchemaParser struct {
	cfg *config.Config
	db  *DB
}

func NewSchemaParser(db *DB, cfg *config.Config) SchemaParser {
	return &defaultSchemaParser{
		cfg: cfg,
		db:  db,
	}
}

func (p *defaultSchemaParser) Parse() (map[string]*Table, error) {
	tables := make(map[string]*Table)
	for tableName := range p.cfg.TableConf {
		columns, err := p.loadColumns(tableName)
		if err != nil {
			return nil, errors.Wrapf(err, "解析表结构失败: %s", tableName)
		}
		for _, column := range columns {
			column.GoType = MapTypeToGo[removeLength(column.Type)]
			column.Lower = utils.ToLowerCamelCase(column.Field)
			column.Upper = utils.ToUpperCamelCase(column.Field)
		}
		rows, err := p.loadRows(tableName)
		if err != nil {
			return nil, errors.Wrapf(err, "加载表数据失败: %s", tableName)
		}

		// 装配Column和Row
		mapColumns := lo.SliceToMap(columns, func(item *Column) (string, *Column) { return item.Field, item })
		for _, row := range rows {
			for _, ele := range row {
				if col, ok := mapColumns[ele.Name]; ok {
					ele.Column = col
					ele.Val, err = p.parseVal(ele)
					if err != nil {
						return nil, errors.Wrapf(err, "解析列数据类型失败: %s", tableName)
					}
				}
			}
		}
		tables[tableName] = &Table{Name: tableName, Columns: columns, Rows: rows}
	}
	return tables, nil
}

func (p *defaultSchemaParser) loadRows(tableName string) ([]Row, error) {
	rowResult := make([]Row, 0)
	rows, err := p.db.Query("SELECT * FROM " + tableName)
	if err != nil {
		return nil, errors.Wrapf(err, "查询表数据失败: %s", tableName)
	}
	defer rows.Close()
	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			return nil, errors.Wrapf(err, "获取列名失败: %s", tableName)
		}
		rawBytesVals := make([][]byte, len(cols))
		vals := make([]interface{}, len(cols))
		for i := range rawBytesVals {
			vals[i] = &rawBytesVals[i]
		}
		err = rows.Scan(vals...)
		if err != nil {
			return nil, errors.Wrapf(err, "扫描行数据失败: %s", tableName)
		}
		row := make([]*Ele, 0)

		for i, val := range vals {
			row = append(row, &Ele{Name: cols[i], Val: val})
		}

		rowResult = append(rowResult, row)
	}
	// 解析表结构
	return rowResult, nil
}

func (p *defaultSchemaParser) loadColumns(tableName string) ([]*Column, error) {
	// 查询表结构
	v := make([]*Column, 0)
	rows, err := p.db.Query("DESCRIBE " + tableName)
	if err != nil {
		return nil, errors.Wrapf(err, "查询表结构失败: %s", tableName)
	}
	defer rows.Close()
	for rows.Next() {
		col := new(Column)
		err = rows.Scan(&col.Field, &col.Type, &col.Null, &col.Key, &col.Default, &col.Extra)
		if err != nil {
			return nil, errors.Wrapf(err, "扫描列数据失败: %s", tableName)
		}
		v = append(v, col)
	}
	return v, nil
}

// 将Ele中的Val值转换成对应的类型
func (p *defaultSchemaParser) parseVal(ele *Ele) (interface{}, error) {
	kind, ok := MapTypeToSQL[removeLength(ele.Column.Type)]
	if !ok {
		return nil, fmt.Errorf("column field %s,  不支持的类型: %s", ele.Column.Field, ele.Column.Type)
	}
	// 将实际类型是[]byte的interface{}转换成对应类型
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.ParseInt(string(*ele.Val.(*[]byte)), 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.ParseUint(string(*ele.Val.(*[]byte)), 10, 64)
	case reflect.Float32, reflect.Float64:
		return strconv.ParseFloat(string(*ele.Val.(*[]byte)), 64)
	case reflect.String:
		return string(*ele.Val.(*[]byte)), nil
	default:
		return nil, fmt.Errorf("column field %s,  不支持的类型: %s", ele.Column.Field, ele.Column.Type)
	}
}
