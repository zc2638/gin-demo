package model

type Module struct {
	AutoID
	Name   string `gorm:"type:varchar(50);unique_index;not null;" json:"name"` // 模块名称
	Mailto string `gorm:"type:varchar(255);" json:"mailto"`                    // 提醒邮件收件人
	Fixed  string `gorm:"type:varchar(50);" json:"fixed"`                      // 提醒周期设置(cron规则)
}
