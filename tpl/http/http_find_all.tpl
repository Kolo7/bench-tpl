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