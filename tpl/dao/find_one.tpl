// {{ .upperTableName}}FindOne 根据 {{ .lowerTableName}}Id 查询 {{ .tableName}} 信息
// error gorm.ErrRecordNotFound 代表未找到记录
func (d *Dao) {{ .upperTableName}}FindOne(ctx context.Context, {{ .lowerTableName}}Id {{.tablePrimaryKey.GoType}}) (*model.{{ .upperTableName}}, error) {
    query := "SELECT * FROM {{ .tableName}} WHERE {{ .tablePrimaryKey.Field}} = ?"
    var {{ .lowerTableName}} model.{{ .upperTableName}}
    if err := ormDB.Raw(query, {{ .lowerTableName}}Id).Scan(&{{ .lowerTableName}}).Error; err!= nil {
        return nil, errors.Join(err, ErrNotFound)
    }
    return &{{ .lowerTableName}}, nil
}