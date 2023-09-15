package main

import (
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/database"
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{ // Let our browser to allow make a request
		AllowCredentials: true, // With this the frontend can get the cookie that we sent and it can send it back
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
