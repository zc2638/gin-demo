package database

import (
	"dc-kanban/lib/logger"
	"dc-kanban/model"
	"github.com/zctod/tool/common/utils"
)

func Seed() {
	seedAdmin()
	seedModule()
	seedConfig()
}

func seedAdmin() {

	db := NewDB()
	admin := model.Admin{
		Name:     "admin",
		Password: utils.MD5("111111"),
	}
	db.First(&admin, admin)
	if admin.ID == 0 {
		db.Create(&admin)
		if db.NewRecord(admin) == true {
			panic("admin数据填充失败")
		}
		logger.Info("admin数据填充成功")
	} else {
		logger.Info("admin数据已填充")
	}
}

func seedModule() {

	var module model.Module

	db := NewDB()
	db.First(&module)
	if module.ID == 0 {
		db.Exec("INSERT INTO `modules` (`name`) VALUES (?), (?), (?), (?), (?), (?), (?), (?);",
			"行业应用项目组", "技术方案与生态合作", "PMO运营", "道客大学", "日常运营", "人事运营", "一周大事记播报", "IT运营")
		logger.Info("module数据填充成功")
	} else {
		logger.Info("module数据已填充")
	}
}

func seedConfig() {

	var config model.Config
	db := NewDB()
	db.First(&config)
	if config.Code == "" {
		db.Exec(
			"INSERT INTO `configs` (`code`, `name`, `value`) VALUES (?), (?), (?), (?)",
			[]string{"mail_from", "发件邮箱", ""},
			[]string{"mail_pwd", "邮箱密码", ""},
			[]string{"mail_subject", "标题", ""},
			[]string{"mail_body", "内容", ""},
		)
		logger.Info("config数据填充成功")
	} else {
		logger.Info("config数据已填充")
	}
}
