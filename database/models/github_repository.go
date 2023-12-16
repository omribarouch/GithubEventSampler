package models

import "gorm.io/gorm"

type GithubRepository struct {
	gorm.Model
	Name string `json:"name" gorm:"text;not null;default:null"`
	Url  string `json:"url" gorm:"url;not null;default:null"`
}

func (GithubRepository) TableName() string {
	return "GithubRepository"
}
