package main

import (
	"fmt"

	db "github.com/ShankaranarayananBR/FiberApp/config"
	routes "github.com/ShankaranarayananBR/FiberApp/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Sales api using go Fiber.....")
	db.Connect()

	// creating fiber app instance
	app := fiber.New()
	app.Use(app)

	// Registering routes in fiber app
	routes.Setup(app)

	// Listening to a port
	app.Listen(":8888")

}
