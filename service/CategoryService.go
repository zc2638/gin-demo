package service

import (
	"dc-kanban/lib/database"
	"dc-kanban/model"
	"errors"
)

type CategoryService struct{ BaseService }

// 模块分类列表
func (s *CategoryService) GetList(page, pageSize string, moduleId interface{}) (cateList []model.ChartCategory, err error) {

	db := database.NewDB()
	if moduleId.(string) != "0" {
		db = db.Where("module_id = ?", moduleId)
	}
	db, err = s.Paginate(db, page, pageSize)
	if err != nil {
		return
	}
	db.Find(&cateList)
	return
}

// 根据模块ID获取模块分类列表
func (s *CategoryService) GetListByModuleIds(moduleIds interface{}) (cateList []model.ChartCategory) {

	db := database.NewDB()
	db.Where("module_id in (?)", moduleIds).Find(&cateList)
	return
}

// 根据ID获取模块分类详情
func (s *CategoryService) GetCategoryByID(id interface{}) (cate model.ChartCategory, err error) {

	db := database.NewDB()

	db.Where("id = ?", id).First(&cate)
	if cate.ID == 0 {
		err = errors.New("不存在的模块分类")
	}
	return
}

// 创建模块分类
func (s *CategoryService) CreateCategory(cate model.ChartCategory) (model.ChartCategory, error) {

	db := database.NewDB()
	db.Create(&cate)
	if db.NewRecord(cate) == true {
		return cate, errors.New("创建失败")
	}
	return cate, nil
}

// 更新模块分类
func (s *CategoryService) UpdateCategory(cate model.ChartCategory) {

	database.NewDB().Save(&cate)
}

// 删除模块分类
func (s *CategoryService) DeleteCategory(cate model.ChartCategory) {

	itemList := new(CategoryItemService).GetListByCateIds(cate.ID)

	db := database.NewDB()
	tx := db.Begin()
	if len(itemList) > 0 {
		itemIds := make([]uint, 0)
		for _, v := range itemList {
			itemIds = append(itemIds, v.ID)
		}
		tx.Where("id in (?)", itemIds).Delete(&model.ChartItem{})
	}
	tx.Delete(&cate)
	tx.Commit()
}