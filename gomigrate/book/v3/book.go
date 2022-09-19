package v3

import (
	"database/sql"
	"sqlmigration/misc"
)

type Book struct {
	Title       string
	Description string
}

const addRows = `
	INSERT INTO books (title, description) VALUES ('Nineteen Eighty-Four', 'authored by George Orwell in 1949');
`

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
