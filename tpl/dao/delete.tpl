// {{.upperTableName}}Delete deletes a {{.tableName}} by ID.
func (d *Dao) {{.upperTableName}}Delete(ctx context.Context, {{.tableUpperPrimaryKeyField}} int64) (err error) {
    db:=d.DB(ctx)
    query := fmt.Sprintf("DELETE FROM {{.tableName}} WHERE {{ .tablePrimaryKey.Field}} = ?")
    if err = db.Exec(query, {{.tableUpperPrimaryKeyField}}).Error; err != nil {
        return err
    }
    return nil
}