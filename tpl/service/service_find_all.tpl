func (s *Service) {{.upperTableName}}GetAll(ctx context.Context, req api.{{.upperTableName}}GetAllReq) (*api.{{.upperTableName}}GetAllResp, ecode.Codes) {
    records, total, err := s.d.GetAll{{.upperTableName}}(ctx, req.PageNum, req.PageSize, req.Order)
    if err!= nil {
        xlog.Error("{{.upperTableName}}GetAll failed: %v", err)
        return nil, ecode.ServerErr
    }

    resp := api.{{.upperTableName}}GetAllResp{
        List: make([]*api.{{.upperTableName}}Resp, 0),
        Total: total,
    }
    for _, record := range records {
        resp.List = append(resp.List, &api.{{.upperTableName}}Resp{
            {{range  $field := .tableColumns}}{{$field.Upper}}: record.{{$field.Upper}},
            {{end}}
        })
    }
    return &resp, ecode.OK
}