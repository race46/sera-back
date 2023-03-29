package middleware

import (
	"github.com/gofiber/fiber/v2"
	"sera-back/utils"
)

func Authenticate(c *fiber.Ctx) error {

	//return c.Next()
	sessionToken := c.Cookies("session_token")
	println(sessionToken)
	println("token")
	user := utils.DecodeJwt(sessionToken)
	if user == "" {
		return c.Status(400).SendString("STOP HACI!!!")
	}
	println(user)
	c.Set("user", "selamke")
	println(c.Get("user"))
	c.Locals("user", user)
	return c.Next()
}
