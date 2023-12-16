package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	index := map[string]string{
		"Application": "Github Event Handler",
		"Author":      "Omri Barouch",
	}
	return c.Status(200).JSON(index)
}
