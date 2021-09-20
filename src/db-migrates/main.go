package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// TODO: iofs has not already released yet
	// _ "github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"local.packages/web"
)

//go:embed migrations/*.sql
var migrations embed.FS

func copyMigrationsToDir(to_dir string) {
	to_dir, _ = filepath.Abs(to_dir)
	sql_files, _ := fs.Glob(migrations, "migrations/*.sql")

	for _, sql_file := range sql_files {
		from_buf, _ := migrations.Open(sql_file)
		defer from_buf.Close()

		to := filepath.Join(to_dir, filepath.Base(sql_file))
		to_buf, _ := os.Create(to)
		defer to_buf.Close()

		_, err := io.Copy(to_buf, from_buf)
		if err != nil {
			log.Fatal("Copy err:", to, err)
		}
	}
}

func main() {
	tmpdir, err := os.MkdirTemp("", "migrations")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	copyMigrationsToDir(tmpdir)

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
	err = m.Steps(2)
	if err != nil {
		log.Fatal(err)
	}
}
