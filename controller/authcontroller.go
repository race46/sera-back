package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"sera-back/database"
	"sera-back/models"
	"sera-back/utils"
	"time"
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
			HTTPOnly: false,
			Secure:   false,
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
	//file := new(models.File)
	username := c.Locals("user").(string)
	database.Connection.Preload("File").First(user, "username=?", username)

	return c.JSON(user)
}

func LogOut(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    "",
		HTTPOnly: false,
		Secure:   false,
		Expires:  time.Now(),
	})

	return c.SendString("ok")
}
