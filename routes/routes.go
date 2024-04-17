package routes

import (
	"github.com/ShankaranarayananBR/FiberApp/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controller.Login)
	app.Get("/cashiers/:cashierId/logout", controller.Logout)
	app.Post("/cashiers:/cashierId/passcode", controller.Passcode)

	// Cashier routes
	app.Post("/cashier", controller.CreateCashier)
	app.Get("/cashiers", controller.GetCashierList)
	app.Put("/cashiers/:cashierId", controller.UpdateCashier)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)
}
