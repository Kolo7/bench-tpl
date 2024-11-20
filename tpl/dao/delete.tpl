// {{.upperTableName}}Delete deletes a {{.tableName}} by ID.
func (d *Dao) {{.upperTableName}}Delete(ctx context.Context, {{.tableUpperPrimaryKeyField}} int64) (err error) {
    query := fmt.Sprintf("DELETE FROM {{.tableName}} WHERE {{ .tablePrimaryKey.Field}} = ?")
    db := ormDB.TxDB(ctx)
    if err = db.Exec(query, {{.tableUpperPrimaryKeyField}}).Error; err != nil {
        return err
    }
    return nil
}