package models

import "gorm.io/gorm"

type GithubEvent struct {
	gorm.Model
}

func (GithubEvent) TableName() string {
	return "GithubEvent"
}
