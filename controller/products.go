package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sera-back/database"
	"sera-back/models"
	"strconv"
)

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}
	user := c.Locals("user").(string)
	product.Username = user
	result := database.Connection.Create(product)

	if result.Error == nil {
		return c.JSON(product)
	}
	log.Print(result.Error)
	return c.Status(400).SendString("error")
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.Connection.Find(&products)
	return c.JSON(products)
}

func MyProducts(c *fiber.Ctx) error {
	var products []models.Product
	user := c.Locals("user").(string)
	database.Connection.Find(&products, "username=?", user)
	return c.JSON(products)
}

func ActivateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	id, _ := strconv.Atoi(c.Params("id"))

	database.Connection.First(product, "id=?", id)
	user := c.Locals("user").(string)
	if user != product.Username {
		return c.Status(400).SendString("not yours")
	}
	product.Active = !product.Active
	database.Connection.Save(product)

	return c.SendString("ok")
}

func DeleteProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	product := new(models.Product)
	id, _ := strconv.Atoi(c.Params("id"))

	result := database.Connection.First(product, "id=?", id)

	if result.Error != nil || product.Username != user {
		return c.Status(400).SendString("error")
	}

	result = database.Connection.Delete(product)

	if result.Error != nil {
		return c.Status(400).SendString("error")
	}

	return c.SendString("ok")

}
