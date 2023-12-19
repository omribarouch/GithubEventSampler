package models

import "time"

type Repository struct {
	ID                       float64   `json:"id" gorm:"column:ID;primaryKey"`
	Name                     string    `json:"name" gorm:"column:Name;text;not null"`
	URL                      string    `json:"url" gorm:"column:Url;url;not null"`
	LastInvolvementTimestamp time.Time `json:"lastInvolvementTimestamp" gorm:"column:LastInvolvementTimestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (Repository) TableName() string {
	return "Repository"
}
