package model

type Admin struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Password string `gorm:"type:varchar(255);not null;" json:"password"`
	Rule     string `gorm:"type:varchar(255);" json:"rule"`
	Timestamps
}
