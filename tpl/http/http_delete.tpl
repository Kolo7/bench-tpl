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