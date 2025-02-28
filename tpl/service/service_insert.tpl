func (s *Service) {{.upperTableName}}Insert(ctx context.Context, req api.{{.upperTableName}}InsertReq)  ecode.Codes {
    err := s.d.{{ .upperTableName}}Insert(ctx, &model.{{.upperTableName}}{
    {{- range $column := .tableColumns}}{{if not (inExcludedFields $column.Upper)}}{{$column.Upper}}: req.{{$column.Upper}}, {{end}}
    {{end}}
    })
    if errors.Is(err, dao.ErrUniqueConflict) {
        return ecode.Transient(ecode.RequestErr, "已存在")
    }else if err != nil {
        xlog.Error("{{.upperTableName}}Insert failed, req: %#v, err: %v", req, err)
        return ecode.ServerErr
    }

    return ecode.OK
}