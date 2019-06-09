package service

import (
	"dc-kanban/lib/database"
	"dc-kanban/model"
	"errors"
	"github.com/zctod/tool/common/utils"
)

type AdminService struct{ BaseService }

// 查询用户列表
func (s *AdminService) GetList(page, pageSize int) (adminList []model.Admin) {

	database.NewDB().Limit(pageSize).Offset((page - 1) * pageSize).Find(&adminList)
	return
}

// 根据ID查询用户
func (s *AdminService) GetAdminByID(id interface{}) (admin model.Admin, err error) {

	database.NewDB().Where("id = ?", id).First(&admin)
	if admin.ID == 0 {
		err = errors.New("不存在的用户")
	}
	return
}

// 根据用户名密码查询用户
func (s *AdminService) GetAdminByName(admin model.Admin) (model.Admin, error) {

	db := database.NewDB()
	if admin.Name != "" {
		db = db.Where("name = ?", admin.Name)
	}
	if admin.Password != "" {
		db = db.Where("password = ?", admin.Password)
	}
	db.First(&admin)
	if admin.ID == 0 {
		return admin, errors.New("用户账号密码错误")
	}
	return admin, nil
}

// 创建用户
func (s *AdminService) CreateAdmin(admin model.Admin) (model.Admin, error) {

	var err error
	db := database.NewDB()
	_, err = s.GetAdminByName(model.Admin{
		Name: admin.Name,
	})
	if err == nil {
		return admin, errors.New("用户已存在")
	}

	admin.Password = utils.MD5(admin.Password)
	db.Create(&admin)
	if db.NewRecord(admin) == true {
		return admin, errors.New("创建失败")
	}

	return admin, nil
}

// 更新用户信息
func (s *AdminService) UpdateAdmin(admin model.Admin) {

	database.NewDB().Save(&admin)
}

// 删除用户信息
func (s *AdminService) DeleteAdmin(admin model.Admin) {

	database.NewDB().Delete(&admin)
}
