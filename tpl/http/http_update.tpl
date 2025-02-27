// @Summary 更新{{.tableName}}记录
// @Description 更新{{.tableName}}记录
// @Tags {{.tableName}}
// @Accept  json
// @Produce  json
// @Param req body api.{{.upperTableName}}UpdateReq true "请求参数"
// @Success 200 {object} http.JSON{Data=api.{{.upperTableName}}UpdateResp} "成功"
// @Router {{.fqdn}}/{{.tableName}}/update [post] // TODO: 生成路由根据业务实际情况修改, 请删除本注释
func {{.upperTableName}}Update(c *gin.Context){
    r := http.NewGin(c)
    req := api.{{.upperTableName}}UpdateReq{}
    if err := c.ShouldBind(&req); err!= nil {
        xlog.Warnc(c, "request {{.upperTableName}}Update error: %v", err)
        JSONRequestError(r, err)
        return
    }
    r.JSON(nil, service.Get().{{.upperTableName}}Update(c.Request.Context(), req))
}