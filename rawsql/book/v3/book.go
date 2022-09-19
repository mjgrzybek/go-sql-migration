package v3

import (
	"database/sql"
	"sqlmigration/misc"
)

type Book struct {
	Title       string
	Description string
}

func MigrateFromV2(db *sql.DB) {
	version_1(db)
	//version_2(db)
}

var migrationQuery1 = []string{
	`ALTER TABLE books ADD COLUMN description TEXT DEFAULT '<unknown>';`,
	`UPDATE books SET description = 'authored by ' || author || ' in ' || IFNULL(year, 0);`,
	`ALTER TABLE books DROP COLUMN author;`,
	`ALTER TABLE books DROP COLUMN year;`,
}

func version_1(db *sql.DB) {
	tx, err := db.Begin()
	misc.PanicOnError(err)

	for _, q := range migrationQuery1 {
		stmt, err := tx.Prepare(q)
		misc.PanicOnError(err)
		defer stmt.Close()
		_, err = stmt.Exec()
		misc.PanicOnError(err)
	}

	err = tx.Commit()
	misc.PanicOnError(err)
}

const migrationQuery2 = `
	BEGIN TRANSACTION;
	ALTER TABLE books ADD COLUMN description TEXT DEFAULT '<unknown>';
	UPDATE books SET description = 'authored by ' || author || ' in ' || IFNULL(year, 0);
	ALTER TABLE books DROP COLUMN author;
	ALTER TABLE books DROP COLUMN year;
	COMMIT;
`

func version_2(db *sql.DB) {
	_, err := db.Exec(migrationQuery2)
	misc.PanicOnError(err)
}

const addRows = `
	INSERT INTO books (title, description) VALUES ('Nineteen Eighty-Four', 'authored by George Orwell in 1949');
`

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
