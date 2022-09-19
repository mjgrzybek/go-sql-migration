package misc

import (
	"database/sql"
	"log"
	"os"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateDb() (*sql.DB, error) {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
