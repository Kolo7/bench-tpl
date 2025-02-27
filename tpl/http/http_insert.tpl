// @Summary 插入{{.tableName}}记录
// @Description 插入{{.tableName}}记录
// @Tags {{.tableName}}
// @Accept  json
// @Produce  json
// @Param req body api.{{.upperTableName}}InsertReq true "请求参数"
// @Success 200 {object} http.JSON{Data=api.{{.upperTableName}}InsertResp} "成功"
// @Router {{.fqdn}}/{{.tableName}}/add [post] // TODO: 生成路由根据业务实际情况修改, 请删除本注释
func {{.upperTableName}}Insert(c *gin.Context){
    r := http.NewGin(c)
    req := api.{{.upperTableName}}InsertReq{}
    if err := c.ShouldBind(&req); err!= nil {
        xlog.Warnc(c, "request {{.upperTableName}}Insert error: %v", err)
        JSONRequestError(r, err)
        return
    }
    r.JSON(nil, service.Get().{{.upperTableName}}Insert(c.Request.Context(), req))
}