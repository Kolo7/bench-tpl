type {{.upperTableName}}GetAllReq struct {
    PageNum  int `form:"pageNum,default=1" binding:"omitempty,gte=1"`
    PageSize int `form:"pageSize,default=10" binding:"omitempty,gte=1,lte=1000"`
    Order    string `form:"order,default={{.tablePrimaryKey.Field}}"`
}

type {{.upperTableName}}GetAllResp struct {
    Total int `json:"total"`
    List []*{{.upperTableName}}Resp `json:"list"`
}