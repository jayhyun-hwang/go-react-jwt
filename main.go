package main

import (
	"go-react-jwt/database"
	"go-react-jwt/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	//only allowed browser
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
