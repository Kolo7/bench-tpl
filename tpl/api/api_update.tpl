type {{.upperTableName}}UpdateReq struct {
    {{.tableUpperPrimaryKeyField}} {{.tablePrimaryKey.GoType}} `json:"{{.tablePrimaryKey.Lower}}" binding:"required,min=0"`
    {{ range $column := .tableColumns }}
    {{ if not (inExcludedFields $column.Upper) -}}
    {{- $column.Upper }} *{{ $column.GoType }} `json:"{{ $column.Lower -}}"` // {{$column.Comment}}
    {{- end -}}
	{{- end}}
}