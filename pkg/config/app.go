package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := os.Getenv("DB_URL")	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// handle error
		panic("failed to connect to database: " + err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
