package handlers

import (
	"GithubEventHandler/ApiServer/database"
	"GithubEventHandler/models"
	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error {
	events := []models.GithubEvent{}
	database.DB.Db.Find(&events)

	return c.Status(200).JSON(events)
}
