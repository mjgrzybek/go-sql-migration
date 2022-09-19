package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"slqmigration/misc"
	v1 "slqmigration/rawsql/book/v1"
	v2 "slqmigration/rawsql/book/v2"
	v3 "slqmigration/rawsql/book/v3"
)

func main() {
	db, err := v1.CreateDb()
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
}

func list1(db *sql.DB) error {
	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func tx1(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
