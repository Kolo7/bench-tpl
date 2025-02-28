func (s *Service) {{.upperTableName}}Update(ctx context.Context, req api.{{.upperTableName}}UpdateReq)  ecode.Codes {
    updated, err := s.d.{{.upperTableName}}FindOne(ctx, req.{{.tableUpperPrimaryKeyField}})
    if errors.Is(err, dao.ErrNotFound) {
        return ecode.RequestErr
    }else if err!= nil {
        xlog.Error("{{.upperTableName}}Update FindOne failed, req: %v, err: %v", req, err)
        return ecode.ServerErr
    }
{{- range $column := .tableColumns}}
    {{if not (inExcludedFields $column.Upper)}}
            if req.{{$column.Upper}} != nil {
                updated.{{$column.Upper}} = *req.{{$column.Upper}}
            }
    {{end}}
{{end}}
    err = s.d.{{ .upperTableName}}Update(ctx, updated)
    if errors.Is(err, dao.ErrUniqueConflict) {
        return ecode.Transient(ecode.RequestErr, "已存在")
    }else if err != nil {
        xlog.Error("{{.upperTableName}}Update failed, req: %v, err: %v", req, err)
        return ecode.ServerErr
    }

    return ecode.OK
}