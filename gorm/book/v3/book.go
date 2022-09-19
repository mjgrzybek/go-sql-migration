package v3

import (
	"gorm.io/gorm"
	"sqlmigration/misc"
	rawsqlV3 "sqlmigration/rawsql/book/v3"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
}

func MigrateFromV2(db *gorm.DB) {
	//migrator := db.Migrator()
	//err := migrator.AddColumn(Book{}, "description")
	//misc.PanicOnError(err)
	// // cannot update column via ORM api

	rawDb, err := db.DB()
	misc.PanicOnError(err)

	rawsqlV3.MigrateFromV2(rawDb) // :(
}
