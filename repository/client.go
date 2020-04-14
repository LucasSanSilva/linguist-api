package repository

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var gormClient *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "user:password@tcp(linguist-mysql-db:3306)/linguist-api?parseTime=True&loc=Local")
	if err != nil {
		print(err.Error())
	}

	db.AutoMigrate(&Lesson{})

	gormClient = db
}

func CloseConnection() {
	gormClient.Close()
}
