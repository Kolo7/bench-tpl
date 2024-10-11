// {{.upperTableName}}Update updates a {{.tableName}} record.
func (d *Dao) {{.upperTableName}}Update(ctx context.Context, record *model.{{.upperTableName}}) (err error) {
    db := d.DB(ctx)
    query := fmt.Sprintf("UPDATE {{.tableName}} SET %s WHERE {{ .tablePrimaryKey.Field}} = ?", {{.lowerTableName}}RowsWithPlaceHolder)
    if err = db.Exec(query, {{range $i, $field := .tableColumnUpperFields}}{{if not (inExcludedFields $field)}}record.{{.}},{{end}}{{end}} record.{{.tableUpperPrimaryKeyField}}).Error; err != nil {
        return err
    }
    return nil
}