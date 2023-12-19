package models

type Actor struct {
	ID           float64 `json:"id" gorm:"column:ID;primaryKey"`
	Login        string  `json:"login" gorm:"column:Login;text;not null"`
	DisplayLogin string  `json:"display_login" gorm:"column:DisplayLogin;text;not null"`
	GravatarID   string  `json:"gravatar_id" gorm:"column:GravatarID;text;not null"`
	URL          string  `json:"url" gorm:"column:Url;url;not null"`
	AvatarURL    string  `json:"avatar_url" gorm:"column:AvatarUrl;url;not null"`
}

func (Actor) TableName() string {
	return "Actor"
}
