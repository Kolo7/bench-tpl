type {{.upperTableName}}Req struct {
    {{.tableUpperPrimaryKeyField}} {{.tablePrimaryKey.GoType}} `form:"{{.tablePrimaryKey.Lower}}" binding:"required,min=0"` // {{.tablePrimaryKey.Comment}}
}

type {{.upperTableName}}Resp struct {
    {{range $column := .tableColumns}}
    {{ if not (eq "Deleted" $column.Upper) -}}
    {{$column.Upper}} {{$column.GoType}} `json:"{{$column.Lower}}"` // {{$column.Comment}}{{end}}{{end}}
}