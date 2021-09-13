package web

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBClientSample() {
	db_config := NewDBConnectConfigFromEnv()
	db, _ := sql.Open("mysql", db_config.DSN())
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Printf("Connected to [%s].\n", version)
}
