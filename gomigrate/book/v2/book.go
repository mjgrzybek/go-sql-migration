package v2

import (
	"database/sql"
	"sqlmigration/misc"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

const addRows = `
	INSERT INTO books (title, author, year) VALUES ('The Little Prince', 'Antoine de Saint-Exupéry', 1943);
`

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
