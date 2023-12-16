package models

import "gorm.io/gorm"

type GithubActor struct {
	gorm.Model
	Name      string `json:"name" gorm:"text;not null;default:null"`
	AvatarUrl string `json:"avatarUrl" gorm:"text;not null;default:null"`
}

func (GithubActor) TableName() string {
	return "GithubActor"
}
