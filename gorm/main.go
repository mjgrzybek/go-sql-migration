package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	v1 "sqlmigration/gorm/book/v1"
	v2 "sqlmigration/gorm/book/v2"
	v3 "sqlmigration/gorm/book/v3"
	"sqlmigration/misc"
)

func main() {
	_ = os.RemoveAll("foo.db")
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create v1
	err = db.AutoMigrate(&v1.Book{})
	misc.PanicOnError(err)
	db.Create(&v1.Book{Title: "The Lord of the Rings"})

	// migrate to v2
	err = db.AutoMigrate(&v2.Book{})
	misc.PanicOnError(err)
	db.Create(&v2.Book{
		Title:  "The Little Prince",
		Author: "Antoine de Saint-Exupéry",
		Year:   1943,
	})

	// migrate to v3
	// !! err = db.AutoMigrate(&v3.Book{})
	v3.MigrateFromV2(db)
	misc.PanicOnError(err)
	db.Create(&v3.Book{
		Title:       "Nineteen Eighty-Four",
		Description: "authored by George Orwell in 1949",
	})

	fmt.Println("")
}
