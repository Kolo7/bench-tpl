func (s *Service) {{.upperTableName}}Get(ctx context.Context, req api.{{.upperTableName}}Req) (*api.{{.upperTableName}}Resp, error) {
    record,err := s.d.{{ .upperTableName}}FindOne(ctx, req.{{.tableUpperPrimaryKeyField}})
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, gorm.ErrRecordNotFound
    }else if err!= nil {
        return nil, err
    }
    resp := api.{{.upperTableName}}Resp{
        {{range  $field := .tableColumns}}{{$field.Upper}}: record.{{$field.Upper}},
        {{end}}
    }
    return &resp, nil
}