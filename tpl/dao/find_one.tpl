// {{ .upperTableName}}FindOne 根据 {{ .lowerTableName}}Id 查询 {{ .tableName}} 信息
// error gorm.ErrRecordNotFound 代表未找到记录
func (d *Dao) {{ .upperTableName}}FindOne(ctx context.Context, {{ .lowerTableName}}Id {{.tablePrimaryKey.GoType}}) (*model.{{ .upperTableName}}, error) {
    query := "SELECT * FROM {{ .tableName}} WHERE {{ .tablePrimaryKey.Field}} = ?"
    row := d.DB(ctx)
    var {{ .lowerTableName}} model.{{ .upperTableName}}
    if err := row.Raw(query, {{ .lowerTableName}}Id).Scan(&{{ .lowerTableName}}).Error; err!= nil {
        return nil, err
    }
    return &{{ .lowerTableName}}, nil
}