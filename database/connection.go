package database

import (
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:1234@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
