package models

import "time"

type Repository struct {
	ID                       float64   `json:"id" gorm:"primaryKey"`
	Name                     string    `json:"name" gorm:"text;not null"`
	URL                      string    `json:"url" gorm:"url;not null"`
	LastInvolvementTimestamp time.Time `json:"lastInvolvementTimestamp" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (Repository) TableName() string {
	return "Repository"
}
