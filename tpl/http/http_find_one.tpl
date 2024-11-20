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