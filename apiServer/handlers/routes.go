package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", Index)
	app.Get("/events", GetEvents)
	app.Get("/repositories/recent", GetRecentRepositories)
	app.Get("/actors/recent", GetRecentActors)
}
