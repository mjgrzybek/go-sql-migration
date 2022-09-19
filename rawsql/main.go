package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sqlmigration/misc"
	v1 "sqlmigration/rawsql/book/v1"
	v2 "sqlmigration/rawsql/book/v2"
	v3 "sqlmigration/rawsql/book/v3"
)

func main() {
	db, err := misc.CreateDb()
	defer db.Close()
	misc.PanicOnError(err)

	// create v1
	v1.CreateV1Table(db)
	v1.AddSampleRows(db)

	// migrate to v2
	v2.MigrateFromV1(db)
	v2.AddSampleRows(db)

	// migrate to v3
	v3.MigrateFromV2(db)
	v3.AddSampleRows(db)

	fmt.Println("")
}
