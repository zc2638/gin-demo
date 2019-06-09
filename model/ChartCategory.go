package model

type ChartCategory struct {
	AutoID
	ModuleId uint   `gorm:"not null;" json:"module_id"`
	Title    string `gorm:"type:varchar(50);not null;" json:"title"`
	Desc     string `gorm:"type:varchar(255)" json:"desc"`
	Timestamps
}
