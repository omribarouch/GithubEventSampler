package models

type GithubActor struct {
	ID        uint          `json:"id" gorm:"column:ID;primaryKey"`
	Name      string        `json:"name" gorm:"column:Name;text;not null;default:null"`
	AvatarUrl string        `json:"avatarUrl" gorm:"column:AvatarUrl;text;not null;default:null"`
	Events    []GithubEvent `gorm:"foreignKey:ActorID"`
}

func (GithubActor) TableName() string {
	return "GithubActor"
}
