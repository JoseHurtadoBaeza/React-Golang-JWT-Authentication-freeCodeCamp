package controllers

import (
	"strconv"
	"time"

	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/database"
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//const SecretKey = "secret"

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

func Login(c *fiber.Ctx) error {

	var data map[string]string // To get the data from the post request

	if err := c.BodyParser(&data); err != nil { // To parse the data
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user) // Get the first result and assign to the user variable

	if user.Id == 0 { // If we don't find the user
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{ // fiber.Map is basically a map[string]interface{}, so we can put anything there
			"message": "User not found",
		})
	}

	// If we got the user, we have to compare the password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// Obtain the current time
	now := time.Now()

	// Add 24 hours to the current time
	expirationTime := now.Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime), // 1 day
		Issuer:    strconv.Itoa(int(user.Id)),         // The issuer is our user
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	mySigningKey := []byte("AllYourBase")

	ss, err := token.SignedString([]byte(mySigningKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// We want to store the signed JWT into a cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    ss,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true, // Because the frontend cannot access this cookie. This cookie is meant only to be stored in frontend and send it, but the frontend doesn't need to access it.
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

func User(c *fiber.Ctx) error {

	mySigningKey := []byte("AllYourBase")

	cookie := c.Cookies("jwt") // Get the cookie

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

// To logout we have to remove the cookie just created before
func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{ // To remove cookies we create another cookie and set the expire time in the past
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expires 1 hour ago
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
