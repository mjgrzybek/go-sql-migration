package v1

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string
}
