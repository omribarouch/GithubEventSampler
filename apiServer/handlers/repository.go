package handlers

import (
	"GithubEventHandler/database"
	"GithubEventHandler/database/models"
	"github.com/gofiber/fiber/v2"
)

func GetRecentRepositories(c *fiber.Ctx) error {
	repositories := []models.Repository{}
	database.DB.Db.Order("last_involvement_timestamp desc").Limit(20).Find(&repositories)

	return c.Status(200).JSON(repositories)
}
