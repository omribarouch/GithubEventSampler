package utils

import (
	"GithubEventHandler/database"
	"GithubEventHandler/database/models"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/segmentio/kafka-go"
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
			newEvent := new(models.GithubEvent)
			err := mapstructure.Decode(rawEvent, &newEvent)
			if err != nil {
				fmt.Println("Error processing message from kafka:", err)
			}

			database.DB.Db.Create(&newEvent)
			fmt.Println("Successfully processed new event:", rawEvent)
		}
	}
}
