
// GetAll{{.upperTableName}} gets all {{.lowerTableName}} records from database.
func (d *Dao) GetAll{{.upperTableName}}(ctx context.Context, pageNum, pageSize int, order string) (records []*model.{{.upperTableName}}, total int64, err error) {
	query := fmt.Sprintf("SELECT * FROM {{.tableName}}")
	searchQuery := fmt.Sprintf("%s ORDER BY %s LIMIT %d OFFSET %d", query, order, pageSize, (pageNum-1)*pageSize)
	
	total, err = Count(ctx, ormDB, query)
	if err != nil  {
		return nil, 0, err
	}
	
	if err = ormDB.Raw(searchQuery).Scan(&records).Error; err != nil  {
		return nil, 0, err
	}
	return records, total, nil
}