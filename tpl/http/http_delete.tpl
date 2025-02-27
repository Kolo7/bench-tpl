// @Summary 删除单个{{.tableName}}
// @Description 删除单个{{.tableName}}
// @Tags {{.tableName}}
// @Accept  json
// @Produce  json
// @Param req body api.{{.upperTableName}}DeleteReq true "请求参数"
// @Success 200 {object} http.JSON{Data=api.{{.upperTableName}}DeleteResp} "成功"
// @Router {{.fqdn}}/{{.tableName}}/delete [post] // TODO: 生成路由根据业务实际情况修改, 请删除本注释
func {{.upperTableName}}Delete(c *gin.Context){
    r := http.NewGin(c)
    req := api.{{.upperTableName}}DeleteReq{}
    if err := c.ShouldBind(&req); err!= nil {
        xlog.Warnc(c, "request {{.upperTableName}}Delete error: %v", err)
        JSONRequestError(r, err)
        return
    }
    r.JSON(nil, service.Get().{{.upperTableName}}Delete(c.Request.Context(), req))
}