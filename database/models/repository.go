package models

type Repository struct {
	ID   float64 `json:"id" gorm:"column:ID;primaryKey"`
	Name string  `json:"name" gorm:"column:Name;text;not null"`
	URL  string  `json:"url" gorm:"column:Url;url;not null"`
}

func (Repository) TableName() string {
	return "Repository"
}
