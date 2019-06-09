package service

import (
	"dc-kanban/lib/database"
	"dc-kanban/model"
	"errors"
)

type NoteService struct{ BaseService }

// 大事记列表
func (s *NoteService) GetList(page, pageSize string, id interface{}) (noteList []model.Note, err error) {

	db := database.NewDB()
	if id.(string) != "0" {
		db = db.Where("module_id = ?", id)
	}

	db, err = s.Paginate(db, page, pageSize)
	if err != nil {
		return
	}

	db.Order("id desc").Find(&noteList)
	return
}

// 根据模块ID获取所有大事记列表
func (s *NoteService) GetAll() (noteList []model.Note) {

	db := database.NewDB()
	db.Find(&noteList)
	return
}

// 根据ID查询单条大事记
func (s *NoteService) GetNoteByID(id interface{}) (note model.Note, err error) {

	database.NewDB().Where("id = ?", id).First(&note)
	if note.ID == 0 {
		err = errors.New("不存在的记录")
	}
	return
}

// 添加大事记
func (s *NoteService) CreateNote(note model.Note) (model.Note, error) {

	db := database.NewDB()
	db.Create(&note)
	if db.NewRecord(note) == true {
		return note, errors.New("创建失败")
	}
	return note, nil
}

// 更新大事记
func (s *NoteService) UpdateNote(note model.Note) int64 {

	db := database.NewDB()
	return db.Save(&note).RowsAffected
}

// 删除大事记
func (s *NoteService) DeleteNote(note model.Note) {

	database.NewDB().Delete(&note)
}
