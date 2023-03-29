package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sera-back/database"
	"sera-back/router"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())

	router.SetUpRoutes(app)

	println(app.Listen(":8080"))
}
