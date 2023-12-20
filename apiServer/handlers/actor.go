package handlers

import (
	"GithubEventHandler/database"
	"GithubEventHandler/database/models"
	"github.com/gofiber/fiber/v2"
)

func GetRecentActors(c *fiber.Ctx) error {
	actors := []models.Actor{}
	database.DB.Db.Order("last_involvement_timestamp desc").Limit(20).Find(&actors)

	return c.Status(200).JSON(actors)
}
