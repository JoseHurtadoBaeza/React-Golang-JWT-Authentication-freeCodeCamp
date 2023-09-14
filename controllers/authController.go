package controllers

import (
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/database"
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string // To get the data from the post request

	if err := c.BodyParser(&data); err != nil { // To parse the data
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) // We hash the password using bcrypt package

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user) // Insert the user in the database

	return c.JSON(user)
}
