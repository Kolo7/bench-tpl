type {{.upperTableName}}GetAllReq struct {
    PageNum  int `form:"pageNum,default=1" binding:"omitempty,gte=1"`   // 页码
    PageSize int `form:"pageSize,default=10" binding:"omitempty,gte=1,lte=1000"` // 每页条数
    Order    string `form:"order,default={{.tablePrimaryKey.Field}}"` // 排序字段
}

type {{.upperTableName}}GetAllResp struct {
    Total int64 `json:"total"` // 总条数
    List []*{{.upperTableName}}Resp `json:"list"`
}