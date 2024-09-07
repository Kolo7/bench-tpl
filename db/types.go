package db

import "reflect"

const (
	TypeDBInt       = "int"
	TypeDBBigInt    = "bigint"
	TypeTinyInt     = "tinyint"
	TypeDBString    = "varchar"
	TypeDBFloat     = "float"
	TypeDBBool      = "bool"
	TypeDBTime      = "datetime"
	TypeDBTimestamp = "timestamp"
)

// sql类型与go类型反射值
var MapTypeToSQL = map[string]reflect.Kind{
	TypeDBInt:    reflect.Int,
	TypeDBBigInt: reflect.Int64,
	TypeDBString: reflect.String,
	TypeDBFloat:  reflect.Float64,
	TypeDBTime:   reflect.String,
	TypeDBBool:   reflect.Bool,
}
