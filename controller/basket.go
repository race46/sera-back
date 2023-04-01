package controller

import (
	"github.com/gofiber/fiber/v2"
	"sera-back/database"
	"sera-back/models"
	"strconv"
)

func AddBasket(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	id, _ := strconv.Atoi(c.Params("id"))
	basket := models.Basket{Username: user, ProductId: int64(id)}
	result := database.Connection.Create(basket)

	if result.Error == nil {
		return c.JSON(basket)
	}
	return c.Status(400).SendString("error")
}

func RemoveBasket(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	id, _ := strconv.Atoi(c.Params("id"))
	//basket := new(models.Basket)

	database.Connection.Delete(&models.Basket{}, "username=? and product_id=?", user, id)

	return c.SendString("ok")
}

func MyBasket(c *fiber.Ctx) error {
	var basket []models.Basket
	user := c.Locals("user").(string)
	database.Connection.Preload("Product").Find(&basket, "username=?", user)
	return c.JSON(basket)
}
