package models

import "time"

type Actor struct {
	ID                       float64   `json:"id" gorm:"primaryKey"`
	Login                    string    `json:"login" gorm:"text;not null"`
	DisplayLogin             string    `json:"display_login" gorm:"text;not null"`
	GravatarID               string    `json:"gravatar_id" gorm:"text;not null"`
	URL                      string    `json:"url" gorm:"url;not null"`
	AvatarURL                string    `json:"avatar_url" gorm:"url;not null"`
	LastInvolvementTimestamp time.Time `json:"lastInvolvementTimestamp" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (Actor) TableName() string {
	return "Actor"
}
