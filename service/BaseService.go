package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
)

type BaseService struct{}

func (s *BaseService) Paginate(db *gorm.DB, page, pageSize string) (*gorm.DB, error) {

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return db, errors.New("请填写page为数值类型")
	}

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		return db, errors.New("请填写pageSize为数值类型")
	}

	db.Limit(pageNum).Offset((pageNum - 1) * pageSizeNum)
	return db, nil
}
