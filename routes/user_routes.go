package routes

import (
	"project/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Post("/", handlers.CreateUser)
}
