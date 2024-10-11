type {{.upperTableName}}InsertReq struct {
	{{- range $column := .tableColumns}}{{if not (inExcludedFields $column.Upper)}}{{$column.Upper}} {{$column.GoType}} `json:"{{$column.Lower}}"`{{end}}
	{{end}}
}