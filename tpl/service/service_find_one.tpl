func (s *Service) {{.upperTableName}}Get(ctx context.Context, req api.{{.upperTableName}}Req) (*api.{{.upperTableName}}Resp, ecode.Codes) {
    record,err := s.d.{{ .upperTableName}}FindOne(ctx, req.{{.tableUpperPrimaryKeyField}})
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ecode.RequestErr
    }else if err!= nil {
        xlog.Error("{{.upperTableName}}FindOne failed, req.{{.tableUpperPrimaryKeyField}}: %v, err: %v", req.{{.tableUpperPrimaryKeyField}}, err)
        return nil, ecode.ServerErr
    }
    resp := api.{{.upperTableName}}Resp{
        {{range  $field := .tableColumns}}{{ if not (eq "Deleted" $field.Upper) -}}{{$field.Upper}}: record.{{$field.Upper}},{{end}}
        {{end}}
    }
    return &resp, ecode.OK
}