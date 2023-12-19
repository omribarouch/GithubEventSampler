package models

type GithubRepository struct {
	ID     uint          `json:"id" gorm:"column:ID;primaryKey"`
	Name   string        `json:"name" gorm:"column:Name;text;not null"`
	Url    string        `json:"url" gorm:"column:Url;url;not null"`
	Events []GithubEvent `gorm:"foreignKey:RepositoryID"`
}

func (GithubRepository) TableName() string {
	return "GithubRepository"
}
