package model

type Note struct {
	AutoID
	ModuleId uint `gorm:"not null;" json:"module_id"`
	Content string `gorm:"type:text;not null;" json:"content"`
	Timestamps
}