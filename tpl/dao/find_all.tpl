
// GetAll{{.upperTableName}} gets all {{.lowerTableName}} records from database.
func (d *Dao) GetAll{{.upperTableName}}(ctx context.Context, pageNum, pageSize int, order string) (records []*model.{{.upperTableName}}, total int64, err error) {
	query := fmt.Sprintf("SELECT * FROM {{.tableName}}")
	searchQuery := fmt.Sprintf("%s ORDER BY ? LIMIT ? OFFSET ?", query)
	db := ormDB.TxDB(ctx)
	total, err = Count(ctx, db, query)
	if err != nil  {
		return nil, 0, err
	}
	
	if err = db.Raw(searchQuery, order, pageSize, (pageNum-1)*pageSize).Scan(&records).Error; err != nil  {
		return nil, 0, err
	}
	return records, total, nil
}