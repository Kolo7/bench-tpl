package db

import "github.com/zeromicro/go-zero/core/stores/sqlx"

type DB struct {
	sqlx.SqlConn
}

func NewDB(dsn string) *DB {
	return &DB{
		SqlConn: sqlx.NewMysql(dsn),
	}
}
