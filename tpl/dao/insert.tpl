func (d *Dao) {{.upperTableName}}Insert(ctx context.Context, record *model.{{.upperTableName}}) (err error) {
    query := fmt.Sprintf("INSERT INTO {{.tableName}} (%s) VALUES (%s)", {{.lowerTableName}}RowsExpectAutoSet, {{.lowerTableName}}AutoPlaceHolder)
    db := d.DB(ctx)
    if err = db.Exec(query, {{range $i, $field := .tableColumnUpperFields}}{{if not (inExcludedFields $field)}}record.{{.}}{{if not (eq $i (sub (len $.tableColumnUpperFields) 1))}}, {{end}}{{end}}{{end}}).Error; err != nil {
        return err
    }
    return nil
}