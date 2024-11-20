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