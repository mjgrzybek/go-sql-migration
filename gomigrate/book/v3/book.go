package v3

import (
	"database/sql"
	"slqmigration/misc"
)

type Book struct {
	Title       string
	Description string
}

const migrationQuery = `BEGIN TRANSACTION;

ALTER TABLE books ADD COLUMN description TEXT DEFAULT '<unknown>';

UPDATE books SET description = 'authored by ' || author || ' in ' || IFNULL(year, 0) ;

ALTER TABLE books DROP COLUMN author;
ALTER TABLE books DROP COLUMN year;

COMMIT
`

const addRows = `
	INSERT INTO books (title, description) VALUES ('Nineteen Eighty-Four', 'authored by George Orwell in 1949');
`

func MigrateFromV2(db *sql.DB) {
	_, err := db.Exec(migrationQuery)
	misc.PanicOnError(err)
}

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
