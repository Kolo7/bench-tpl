type {{.upperTableName}}Req struct {
    {{.tableUpperPrimaryKeyField}} {{.tablePrimaryKey.GoType}} `form:"{{.tablePrimaryKey.Lower}}" binding:"required,min=0"`
}

type {{.upperTableName}}Resp struct {
    {{range $column := .tableColumns}}{{$column.Upper}} {{$column.GoType}} `json:"{{$column.Lower}}"`
    {{end}}
}