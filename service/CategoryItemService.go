package service

import (
	"dc-kanban/lib/database"
	"dc-kanban/model"
	"errors"
)

type CategoryItemService struct{ BaseService }

// 根据模块分类ID获取项目列表
func (s *CategoryItemService) GetListByCateID(categoryId interface{}, page, pageSize string) (itemList []model.ChartItem, err error) {

	db := database.NewDB()
	db = db.Where("category_id = ?", categoryId)
	db, err = s.Paginate(db, page, pageSize)
	if err != nil {
		return
	}
	db.Find(&itemList)
	return
}

// 查询指定分类下的所有项目
func (s *CategoryItemService) GetListByCateIds(categoryIds interface{}) (itemList []model.ChartItem) {

	db := database.NewDB()
	db.Where("category_id in (?)", categoryIds).Order("sort desc").Find(&itemList)
	return
}

// 根据ID查询详情
func (s *CategoryItemService) GetItemByID(id interface{}) (item model.ChartItem) {

	db := database.NewDB()
	db.Where("id = ?", id).First(&item)
	return
}

// 创建项目
func (s *CategoryItemService) CreateItem(item model.ChartItem) (model.ChartItem, error) {

	db := database.NewDB()
	db.Create(&item)
	if db.NewRecord(item) == true {
		return item, errors.New("创建失败")
	}
	return item, nil
}

// 更新项目
func (s *CategoryItemService) UpdateItem(item model.ChartItem) int64 {

	db := database.NewDB()
	return db.Save(&item).RowsAffected
}

// 删除项目
func (s *CategoryItemService) DeleteItem(item model.ChartItem) {

	database.NewDB().Delete(&item)
}
