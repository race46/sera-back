package controller

import (
	"github.com/gofiber/fiber/v2"
	"sera-back/database"
	"sera-back/models"
)

func GetSellers(c *fiber.Ctx) error {
	var sellers []models.Product
	database.Connection.Find(&sellers)
	return c.JSON(sellers)
}
