package utils

import (
	"GithubEventHandler/database"
	"GithubEventHandler/database/models"
	"GithubEventHandler/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"time"
)

func ProcessEvents(ch chan kafka.Message) {
	for kafkaMessage := range ch {
		var rawEvents []map[string]interface{}

		err := json.Unmarshal(kafkaMessage.Value, &rawEvents)
		if err != nil {
			fmt.Println("Error parsing the message from kafka:", err)
			continue
		}
		for _, rawEvent := range rawEvents {
			err, newEvent := Deserialize(rawEvent)
			if err != nil {
				fmt.Println("Error processing message from kafka:", err)
			}
			tx := database.DB.Db.Begin()
			EnrichEventData(newEvent)
			database.DB.Db.Create(&newEvent)
			newEvent.Repository.LastInvolvementTimestamp = time.Now()
			newEvent.Actor.LastInvolvementTimestamp = time.Now()
			tx.Commit()
			fmt.Println("Successfully processed new event:", rawEvent)
		}
	}
}

func EnrichEventData(event *models.Event) {
	var repo models.Repository

	err := database.DB.Db.First(&repo, event.RepositoryID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		res := utils.SendGithubRequest("GET",
			event.Repository.URL)

		var enrichedRepositoryData map[string]interface{}

		err := json.Unmarshal(res, &enrichedRepositoryData)
		if err != nil {
			fmt.Println("Error parsing the message from kafka:", err)
		}

		event.Repository.Stars = enrichedRepositoryData["stargazers_count"].(float64)
	}
}

func Deserialize(rawEvent map[string]interface{}) (error, *models.Event) {
	actorData, ok := rawEvent["actor"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("missing or invalid 'actor' field"), nil
	}

	actor := models.Actor{
		ID:           actorData["id"].(float64),
		Login:        actorData["login"].(string),
		DisplayLogin: actorData["display_login"].(string),
		GravatarID:   actorData["gravatar_id"].(string),
		URL:          actorData["url"].(string),
		AvatarURL:    actorData["avatar_url"].(string),
	}

	repoData, ok := rawEvent["repo"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("missing or invalid 'repo' field"), nil
	}

	repo := models.Repository{
		ID:   repoData["id"].(float64),
		Name: repoData["name"].(string),
		URL:  repoData["url"].(string),
	}

	event := &models.Event{
		ID:         rawEvent["id"].(string),
		Type:       rawEvent["type"].(string),
		Actor:      actor,
		Repository: repo,
		CreatedAt:  parseTime(rawEvent["created_at"].(string)),
	}

	return nil, event
}

func parseTime(input string) time.Time {
	layout := "2006-01-02T15:04:05Z"
	t, _ := time.Parse(layout, input)
	return t
}
