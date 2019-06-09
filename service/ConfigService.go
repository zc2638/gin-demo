package service

import (
	"dc-kanban/lib/database"
	"dc-kanban/model"
)

type ConfigService struct{ BaseService }

// 根据标识获取配置列表
func (s *ConfigService) GetListByCodes(codes []string) (configList []model.Config) {

	db := database.NewDB()
	db.Where("code in (?)", codes).Find(&configList)
	return
}

// 根据标识获取单个配置
func (s *ConfigService) GetConfigByCode(code string) (config model.Config) {

	db := database.NewDB()
	db.Where("code = ?", code).First(&config)
	return
}

// 更新单个配置
func (s *ConfigService) UpdateConfig(config model.Config) int64 {

	db := database.NewDB()
	return db.Save(&config).RowsAffected
}

// 批量更新
func (s *ConfigService) BatchUpdate(configs []model.Config) {

	db := database.NewDB()
	tx := db.Begin()
	for _, v := range configs {
		tx.Model(&v).Where("code = ?", v.Code).Update("value", v.Value)
	}
	tx.Commit()
}