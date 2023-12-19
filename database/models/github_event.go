package models

import "time"

type GithubEvent struct {
	ID           uint      `gorm:"column:ID;primaryKey"`
	Timestamp    time.Time `json:"timestamp" gorm:"column:Timestamp;not null;index"`
	Type         string    `json:"type" gorm:"column:Type;text;not null;index"`
	ActorID      uint      `json:"actorId" gorm:"column:ActorID;index"`
	RepositoryID uint      `json:"repositoryId" gorm:"column:RepositoryID;index"`
	Actor        GithubActor
	Repository   GithubRepository
}

func (GithubEvent) TableName() string {
	return "GithubEvent"
}
