package web

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBClient() {
	db, _ := sql.Open("mysql", "app:app@tcp(db:3306)/app_db")
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Printf("Connected to [%s].\n", version)
}
