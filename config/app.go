package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/gobookms?charset=utf8mb4&parseTime=True&loc=Local"	
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
