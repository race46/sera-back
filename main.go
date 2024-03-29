package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"sera-back/database"
	"sera-back/router"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/files", "./files")

	router.SetUpRoutes(app)

	log.Fatalln(app.Listen(":8080"))
}
