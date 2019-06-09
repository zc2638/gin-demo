package database

import (
	"dc-kanban/config"
	"dc-kanban/lib/logger"
	"dc-kanban/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func init() {
	cfg := config.Cfg
	optionStr := cfg.SqlUsername + ":" + cfg.SqlPassword + "@tcp(" + cfg.SqlHost + ":" + cfg.SqlPort + ")/" + cfg.SqlDb + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open("mysql", optionStr)
	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10) // 连接池的空闲数大小
	DB.DB().SetMaxOpenConns(100) // 最大打开连接数
	DB.DB().SetConnMaxLifetime(time.Hour * 6) // 连接最长存活时间

	DBMigrate()
}

func NewDB() *gorm.DB {
	return DB
}

func DBMigrate() {

	// 禁用表名复数
	//DB.SingularTable(true)

	// 自动生成表结构
	DB.AutoMigrate(
		&model.Admin{},
		&model.Module{},
		&model.ChartCategory{},
		&model.ChartItem{},
		&model.Note{},
		&model.Config{},
	)
	logger.Info("数据表迁移成功！")

	Seed()
}