func (d *Dao) {{.upperTableName}}Insert(ctx context.Context, record *model.{{.upperTableName}}) (err error) {
    query := fmt.Sprintf("INSERT INTO {{.tableName}} (%s) VALUES (%s)", {{.lowerTableName}}RowsExpectAutoSet, {{.lowerTableName}}AutoPlaceHolder)
    if err = ormDB.Exec(query, {{range $i, $field := .tableColumnUpperFields}}{{if not (inExcludedFields $field)}}record.{{.}}{{if not (eq $i (sub (len $.tableColumnUpperFields) 1))}}, {{end}}{{end}}{{end}}).Error; err != nil {
        mysqlErr := &mysql.MySQLError{}
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.Join(err, ErrUniqueConflict)
		}
        return err
    }
    return nil
}