// @Summary 获取{{.tableName}}详情
// @Description 获取{{.tableName}}详情
// @Tags {{.tableName}}
// @Produce  json
// @Param req query api.{{.upperTableName}}Req true "请求参数"
// @Success 200 {object} http.JSON{Data=api.{{.upperTableName}}Resp} "成功"
// @Router {{.fqdn}}/{{.tableName}}/get [get] // TODO: 生成路由根据业务实际情况修改, 请删除本注释
func {{.upperTableName}}Get(c *gin.Context){
    r := http.NewGin(c)
    req := api.{{.upperTableName}}Req{}
    if err := c.ShouldBind(&req); err!= nil {
        xlog.Warnc(c, "request {{.upperTableName}}Get error: %v", err)
        JSONRequestError(r, err)
        return
    }
    r.JSON(service.Get().{{.upperTableName}}Get(c.Request.Context(), req))
}