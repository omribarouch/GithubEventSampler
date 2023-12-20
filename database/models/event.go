package models

import (
	"time"
)

type Event struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"not null;default:CURRENT_TIMESTAMP;index"`
	Type         string     `json:"type" gorm:"text;not null;index"`
	ActorID      float64    `json:"actorId"`
	RepositoryID float64    `json:"repositoryId"`
	Actor        Actor      `json:"actor"`
	Repository   Repository `json:"repo"`
}

func (Event) TableName() string {
	return "Event"
}
