package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	v2 "sqlmigration/gomigrate/book/v2"
	v3 "sqlmigration/gomigrate/book/v3"
	"sqlmigration/misc"
	v1 "sqlmigration/rawsql/book/v1"
)

func main() {
	db, err := misc.CreateDb()
	defer db.Close()
	misc.PanicOnError(err)

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	misc.PanicOnError(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://gomigrate/migrations/",
		"foo",
		driver,
	)
	misc.PanicOnError(err)

	// create v1
	err = m.Steps(1)
	misc.PanicOnError(err)
	v1.AddSampleRows(db)

	// migrate to v2
	err = m.Steps(1)
	misc.PanicOnError(err)
	v2.AddSampleRows(db)

	// migrate to v3
	err = m.Steps(1)
	misc.PanicOnError(err)
	v3.AddSampleRows(db)

	fmt.Println("")
}
