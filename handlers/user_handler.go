package handlers

import (
	"log"
	"project/models"
	"project/services"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Print("error on CreateUser bodyparser:", err)
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}
	services.CreateUser(user)
	log.Printf("User created: %v", user)
	return c.JSON(user)
}
