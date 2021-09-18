// //go:build tools

package main

import (
	"embed"
	// 2021/09/18 iofs has not already released yet
	//	_ "github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

func main() {
	//iofs.New(fs, "migrations")
	// TODO: maybe should move to another module (neither web and HERE)
	//	db := web.NewDBClient()
	//	driver, err := mysql.WithInstance(db, &mysql.Config{})
	//    m, err := migrate.NewWithDatabaseInstance(
	//        "file:///migrations",
	//        "mysql", driver)
	//    )
	//    m.Steps(2)
}
