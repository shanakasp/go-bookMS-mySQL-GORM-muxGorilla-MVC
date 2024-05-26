package models

import (
	"github.com/princesp/go-bookms/pkg/config"

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
