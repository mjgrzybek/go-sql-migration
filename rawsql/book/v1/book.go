package v1

import (
	"database/sql"
	"sqlmigration/misc"
)

const createTable = `
	CREATE TABLE IF NOT EXISTS books (id INTEGER NOT NULL PRIMARY KEY, title TEXT);
`

var addRows = `
	INSERT INTO books (title) VALUES ("The Lord of the Rings");
`

type Book struct {
	Title string
}

func CreateV1Table(db *sql.DB) {
	_, err := db.Exec(createTable)
	misc.PanicOnError(err)
}

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
