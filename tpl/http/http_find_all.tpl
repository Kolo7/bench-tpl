// @Summary {{.tableName}}列表
// @Description {{.tableName}}列表
// @Tags {{.tableName}}
// @Produce  json
// @Param req query api.{{.upperTableName}}GetAllReq true "请求参数"
// @Success 200 {object} http.JSON{Data=api.{{.upperTableName}}GetAllResp} "成功"
// @Router {{.fqdn}}/{{.tableName}}/list [get] // TODO: 生成路由根据业务实际情况修改, 请删除本注释
func {{.upperTableName}}GetAll(c *gin.Context){
    r := http.NewGin(c)
    req := api.{{.upperTableName}}GetAllReq{}
    if err := c.ShouldBind(&req); err!= nil {
        xlog.Warnc(c, "request {{.upperTableName}}GetAllReq error: %v", err)
        JSONRequestError(r, err)
        return
    }
    r.JSON(service.Get().{{.upperTableName}}GetAll(c.Request.Context(), req))
}