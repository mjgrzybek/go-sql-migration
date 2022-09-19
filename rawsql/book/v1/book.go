package v1

import (
	"database/sql"
	"log"
	"os"
	"slqmigration/misc"
)

const createTable = `
	CREATE TABLE IF NOT EXISTS books (id INTEGER NOT NULL PRIMARY KEY, title TEXT);
`

const addRows = `
	INSERT INTO books (title) VALUES 
	  ("The Lord of the Rings"); 
-- 	  ("The Little Prince"), 
-- 	  ("Nineteen Eighty-Four");
`

type Book struct {
	Title string
}

func CreateV1Table(db *sql.DB) {
	_, err := db.Exec(createTable)
	misc.PanicOnError(err)
}

func CreateDb() (*sql.DB, error) {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
