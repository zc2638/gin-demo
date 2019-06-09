package service

import (
	"dc-kanban/lib/database"
	"dc-kanban/model"
	"errors"
)

type ModuleService struct{ BaseService }

// 模块列表
func (s *ModuleService) GetAll() (moduleList []model.Module) {

	db := database.NewDB()
	db.Order("id").Find(&moduleList)
	return
}

// 根据ID查询模块
func (s *ModuleService) GetModuleByID(id interface{}) (module model.Module, err error) {

	db := database.NewDB()
	db.Where("id = ?",  id).First(&module)
	if module.ID == 0 {
		err = errors.New("不存在的模块")
	}
	return
}

// 根据ID查询模块列表
func (s *ModuleService) GetModuleByIds(ids interface{}) (moduleList []model.Module) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Order("id").Find(&moduleList)
	return
}

// 更新模块
func (s *ModuleService) UpdateModule(module model.Module) int64 {

	db := database.NewDB()
	return db.Save(&module).RowsAffected
}