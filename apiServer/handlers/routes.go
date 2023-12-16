package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", Index)
	app.Get("/events", GetEvents)
}
