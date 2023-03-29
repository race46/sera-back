package router

import (
	"github.com/gofiber/fiber/v2"
	"sera-back/controller"
	"sera-back/gdive"
	"sera-back/middleware"
)

func SetUpRoutes(app *fiber.App) {
	//Authentication
	app.Post("/api/upload", gdive.UploadFile)
	app.Post("/api/signup", controller.Signup)
	app.Post("/api/login", controller.Login)

	//products
	app.Get("/api/products", controller.GetProducts)

	app.Use(middleware.Authenticate)

	//Authentication
	app.Get("/api/me", controller.Me)

	//Products
	app.Post("/api/product", controller.AddProduct)
	app.Get("/api/my-products", controller.MyProducts)
	app.Put("/api/activate-product/:id", controller.ActivateProduct)

}
