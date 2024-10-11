// Code generated by goctl. DO NOT EDIT.

package {{.lowerPkgName}}

import (
    "context"
    "errors"
    "fmt"
    "strings"

    "{{.modelPackageName}}"
)

var (
    {{.lowerTableName}}FieldNames = RawFieldNames(&model.{{.upperTableName}}{})
    {{.lowerTableName}}Rows       = strings.Join({{.lowerTableName}}FieldNames, ",")
    {{.lowerTableName}}RowsExpectAutoSet   = strings.Join(Remove({{.lowerTableName}}FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`", "`deleted`"), ",")
    {{.lowerTableName}}AutoPlaceHolder = placeHolder(len(Remove({{.lowerTableName}}FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`", "`deleted`")))
	{{.lowerTableName}}RowsWithPlaceHolder = strings.Join(Remove({{.lowerTableName}}FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`", "`deleted`"), "=?,") + "=?"
)

{{template "find_all.tpl" .}}

{{template "find_one.tpl" .}}

{{template "insert.tpl" .}}

{{template "update.tpl" .}}

{{template "delete.tpl" .}}