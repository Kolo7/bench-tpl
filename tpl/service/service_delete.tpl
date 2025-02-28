func (s *Service) {{.upperTableName}}Delete(ctx context.Context, req api.{{.upperTableName}}DeleteReq)  error {
    _, err := s.d.{{.upperTableName}}FindOne(ctx, req.{{.tableUpperPrimaryKeyField}})
    if errors.Is(err, dao.ErrNotFound) {
        return ecode.RequestErr
    }else if err!= nil {
        xlog.Error("{{.upperTableName}}Delete FindOne failed, req: %v, err: %v", req, err)
        return ecode.ServerErr
    }

    if err = s.d.{{ .upperTableName}}Delete(ctx, req.{{.tableUpperPrimaryKeyField}}); err != nil {
        xlog.Error("{{.upperTableName}}Delete failed, req: %v, err: %v", req, err)
        return ecode.ServerErr
    }

    return ecode.OK
}