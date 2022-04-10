package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBClient(config *DBConnectConfig) *sql.DB {
	dsn := config.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
