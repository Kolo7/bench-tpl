package db

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DB struct {
	*sql.DB
}

func NewDB(dsn string) *DB {
	db, err := sqlx.NewMysql(dsn).RawDB()
	if err != nil {
		panic(err)
	}
	return &DB{
		DB: db,
	}
}
