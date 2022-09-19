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

const migrationQuery = `BEGIN TRANSACTION;
ALTER TABLE books ADD COLUMN author TEXT DEFAULT '<unknown>';
ALTER TABLE books ADD COLUMN year INTEGER;
COMMIT
`

const addRows = `
	INSERT INTO books (title, author, year) VALUES ('The Little Prince', 'Antoine de Saint-Exupéry', 1943);
`

func MigrateFromV1(db *sql.DB) {
	_, err := db.Exec(migrationQuery)
	misc.PanicOnError(err)
}

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
