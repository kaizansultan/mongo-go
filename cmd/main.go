package main

import (
	"log"

	"project/database"
	// "project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Server is setting up")
	database.ConnectDB()

	app := fiber.New()

	// routes.RegisterUserRoutes(app)

	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
