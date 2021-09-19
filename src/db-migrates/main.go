// //go:build tools

package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// TODO: iofs has not already released yet
	// _ "github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"local.packages/web"
)

//go:embed migrations/*.sql
var fs embed.FS

func main() {

	var tmpdir string
	tmpdir, err := os.MkdirTemp("", "migrations")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// TODO: Write embed.FS to tempdir

	// TODO: maybe should move to another module (neither web and HERE)
	db_config := web.NewDBConnectConfigFromEnv()
	db := web.NewDBClient(db_config)
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", tmpdir),
		"mysql", driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	m.Steps(2)
}
