package database

import (
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB_URI string = "root@tcp(localhost:3306)/test"

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open(DB_URI), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	DB.AutoMigrate(models.User{})
}
