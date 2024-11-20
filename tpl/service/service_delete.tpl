func (s *Service) {{.upperTableName}}Delete(ctx context.Context, req api.{{.upperTableName}}DeleteReq)  error {
    _, err := s.dao.{{.upperTableName}}FindOne(ctx, req.{{.tableUpperPrimaryKeyField}})
    if errors.Is(err, dao.ErrNotFound) {
        return err
    }else if err!= nil {
        xlog.Error("{{.upperTableName}}Delete FindOne failed, req: %v, err: %v", req, err)
        return err
    }

    if err = s.dao.{{ .upperTableName}}Delete(ctx, req.{{.tableUpperPrimaryKeyField}}); err != nil {
        xlog.Error("{{.upperTableName}}Delete failed, req: %v, err: %v", req, err)
        return err
    }

    return nil
}