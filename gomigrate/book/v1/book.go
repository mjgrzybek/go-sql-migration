package v1

import (
	"database/sql"
	"sqlmigration/misc"
	"strconv"
	"time"
)

type Book struct {
	Title string
}

var addRows = `
	INSERT INTO books (title) VALUES 
	  ("The Lord of the Rings"), 
	  ("` + strconv.FormatInt(time.Now().UnixMilli()%1000, 10) + `");`

func AddSampleRows(db *sql.DB) {
	_, err := db.Exec(addRows)
	misc.PanicOnError(err)
}
