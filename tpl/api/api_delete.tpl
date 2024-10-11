type {{.upperTableName}}DeleteReq struct {
    {{.tableUpperPrimaryKeyField}} {{.tablePrimaryKey.GoType}} `form:"{{.tablePrimaryKey.Lower}}" binding:"required,min=0"`
}