package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/lacnguyen/go-jwt/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gojwt"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to DB!")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
