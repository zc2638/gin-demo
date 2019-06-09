package model

import (
	"time"
)

type AutoID struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
}

type Timestamps struct {
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

