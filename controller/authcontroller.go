package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"sera-back/database"
	"sera-back/models"
	"sera-back/utils"
)

func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	result := database.Connection.Create(user)
	if result.Error == nil {
		jwt := utils.CreateJwt(user)
		c.Cookie(&fiber.Cookie{
			Name:     "session_token",
			Value:    jwt,
			HTTPOnly: true,
			Secure:   true,
		})

		return c.JSON(user)
	}
	return c.Status(400).SendString("error")
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	userDb := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	result := database.Connection.First(&userDb, "username=?", user.Username)
	err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))
	if err != nil || result.Error != nil {
		fmt.Println("Invalid password")
		return c.Status(400).SendString("error")
	} else {
		jwt := utils.CreateJwt(user)
		c.Cookie(&fiber.Cookie{
			Name:     "session_token",
			Value:    jwt,
			HTTPOnly: false,
			Secure:   false,
		})

		return c.JSON(user)
	}

}

func Me(c *fiber.Ctx) error {
	user := new(models.User)
	username := c.Locals("user").(string)
	println(username + " username xdxd")
	database.Connection.First(user, "username=?", username)

	return c.JSON(user)
}
