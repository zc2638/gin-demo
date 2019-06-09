package model

type Config struct {
	Code  string `gorm:"type:varchar(50);not null;primary_key" json:"code"`
	Name  string `gorm:"type:varchar(50);not null;" json:"name"`
	Value string `gorm:"type:text;" json:"value"`
}
