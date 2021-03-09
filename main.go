package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
)

func main() {
	// connect to database
	database.Connect()

	// initialize
	app := fiber.New()

	// call a routes
	routes.Setup(app)

	// listen port
	app.Listen(":8000")
}
