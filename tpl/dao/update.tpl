// {{.upperTableName}}Update updates a {{.tableName}} record.
func (d *Dao) {{.upperTableName}}Update(ctx context.Context, record *model.{{.upperTableName}}) (err error) {
    query := fmt.Sprintf("UPDATE {{.tableName}} SET %s WHERE {{ .tablePrimaryKey.Field}} = ?", {{.lowerTableName}}RowsWithPlaceHolder)
    if err = ormDB.Exec(query, {{range $i, $field := .tableColumnUpperFields}}{{if not (inExcludedFields $field)}}record.{{.}},{{end}}{{end}} record.{{.tableUpperPrimaryKeyField}}).Error; err != nil {
        mysqlErr := &mysql.MySQLError{}
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.Join(err, ErrUniqueConflict)
		}
        return err
    }
    return nil
}