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
	app.Get("/api/logout", controller.LogOut)

	//products
	app.Get("/api/products", controller.GetProducts)

	//sellers
	app.Get("/api/sellers", controller.GetSellers)

	app.Use(middleware.Authenticate)

	//Authentication
	app.Get("/api/me", controller.Me)

	//User
	app.Post("/api/update-pp", gdive.UploadFile, controller.UpdatePP)

	//Products
	app.Post("/api/product", controller.AddProduct)
	app.Get("/api/my-products", controller.MyProducts)
	app.Put("/api/activate-product/:id", controller.ActivateProduct)
	app.Delete("/api/product/:id", controller.DeleteProduct)

	//Basket
	app.Post("/api/basket/:id", controller.AddBasket)
	app.Delete("/api/basket/:id", controller.RemoveBasket)
	app.Get("/api/my-basket", controller.MyBasket)

}
