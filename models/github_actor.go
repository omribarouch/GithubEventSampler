package models

import "gorm.io/gorm"

type GithubActor struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"text;not null;default:null"`
	AvatarUrl string `json:"avatarUrl" gorm:"text;not null;default:null"`
}
