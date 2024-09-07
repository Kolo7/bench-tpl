{{range $index, $value := .db.sys_config.Columns}}
index:{{$index}},value:{{$value.Field}}
{{end}}