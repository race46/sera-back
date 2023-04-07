package middleware

import (
	"github.com/gofiber/fiber/v2"
	"sera-back/utils"
)

func Authenticate(c *fiber.Ctx) error {

	sessionToken := c.Cookies("session_token")

	user := utils.DecodeJwt(sessionToken)
	if user == "" {
		return c.Status(400).SendString("STOP HACI!!!")
	}

	c.Locals("user", user)
	return c.Next()
}
