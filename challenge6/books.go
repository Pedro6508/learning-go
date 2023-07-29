package library

import (
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
}

type Books struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}
