package main

import (
	"GithubEventHandler/apiServer/handlers"
	"GithubEventHandler/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	handlers.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error serving api:", err)
	}
}
