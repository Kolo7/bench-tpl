type {{.upperTableName}}DeleteReq struct {
    {{.tableUpperPrimaryKeyField}} {{.tablePrimaryKey.GoType}} `json:"{{.tablePrimaryKey.Lower}}" binding:"required,min=0"` // {{.tablePrimaryKey.Comment}}
}