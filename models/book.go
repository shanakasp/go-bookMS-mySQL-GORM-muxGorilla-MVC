package models

import (
	"github.com/princesp/go-bookms/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Publication string
}

func init() {
	config.ConnectDb()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

// func (b *Book) CreateBook() *Book {
    
// 	DB.NewRecord(b)
// 	DB.Create(&b)
// 	return b
// }