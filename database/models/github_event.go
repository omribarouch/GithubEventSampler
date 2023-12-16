package models

type GithubEvent struct {
	Type string `json:"type" gorm:"text;not null;default:null"`
}

func (GithubEvent) TableName() string {
	return "GithubEvent"
}
