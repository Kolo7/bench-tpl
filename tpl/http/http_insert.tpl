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