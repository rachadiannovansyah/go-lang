package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	// connect to database
	database.Connect()

	// initialize
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// call a routes
	routes.Setup(app)

	// listen port
	app.Listen(":8000")
}
