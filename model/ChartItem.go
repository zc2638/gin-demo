package model

type ChartItem struct {
	AutoID
	CategoryId uint `gorm:"not null;" json:"category_id"`
	Type uint `gorm:"type:tinyint;not null;" json:"type"`
	Title string `gorm:"type:varchar(50);" json:"title"`
	Content string `gorm:"type:varchar(255)" json:"content"`
	Value string `json:"value"`
	Sort uint `gorm:"type:int;default:1;not null;" json:"sort"`
	Timestamps
}