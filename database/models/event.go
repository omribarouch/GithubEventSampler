package models

import (
	"time"
)

type Event struct {
	ID           string     `json:"id" gorm:"column:ID;primaryKey"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:CreatedAt;not null;default:CURRENT_TIMESTAMP;index"`
	Type         string     `json:"type" gorm:"column:Type;text;not null;index"`
	ActorID      float64    `json:"actorId" gorm:"column:ActorID"`
	RepositoryID float64    `json:"repositoryId" gorm:"column:RepositoryID"`
	Actor        Actor      `json:"actor"`
	Repository   Repository `json:"repo"`
}

func (Event) TableName() string {
	return "Event"
}
