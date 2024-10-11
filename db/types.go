package db

import (
	"reflect"
	"regexp"
)

const (
	TypeDBInt       = "int"
	TypeDBBigInt    = "bigint"
	TypeTinyInt     = "tinyint"
	TypeDBString    = "varchar"
	TypeDBFloat     = "float"
	TypeDBBool      = "bool"
	TypeDBTime      = "datetime"
	TypeDBTimestamp = "timestamp"
	TypeText        = "text"
)

// sql类型与go类型反射值
var MapTypeToSQL = map[string]reflect.Kind{
	TypeDBInt:       reflect.Int,
	TypeDBBigInt:    reflect.Int64,
	TypeDBString:    reflect.String,
	TypeDBFloat:     reflect.Float64,
	TypeDBTime:      reflect.String,
	TypeDBBool:      reflect.Bool,
	TypeDBTimestamp: reflect.String,
	TypeTinyInt:     reflect.Int,
	TypeText:        reflect.String,
}

var MapTypeToGo = map[string]string{
	TypeDBInt:       "int",
	TypeDBBigInt:    "int64",
	TypeDBString:    "string",
	TypeDBFloat:     "float64",
	TypeDBTime:      "time.Time",
	TypeDBBool:      "bool",
	TypeDBTimestamp: "time.Time",
	TypeTinyInt:     "int",
	TypeText:        "string",
}

// 将可变长度的sql类型去除长度
func removeLength(t string) string {
	// 匹配字符串(12) (12,12)
	reg := `\(\d+(,\d+)?\)`
	res := regexp.MustCompile(reg)
	t = res.ReplaceAllString(t, "")
	return t
}
