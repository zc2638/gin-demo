package admin

import (
	"dc-kanban/controller"
	"dc-kanban/model"
	"dc-kanban/service"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
)

type NoteController struct{ controller.Base }

func (t *NoteController) List(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "15")
	moduleId := c.DefaultQuery("moduleId", "0")

	noteList, err := new(service.NoteService).GetList(page, pageSize, moduleId)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	data, err := utils.ArrayStructToMap(noteList)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	t.Data(c, data)
}

func (t *NoteController) Create(c *gin.Context) {

	_ = c.Request.ParseForm()

	moduleId := c.DefaultPostForm("moduleId", "0")
	content := c.PostForm("content")
	if content == "" {
		t.Err(c, "内容不能为空")
		return
	}

	module, err := new(service.ModuleService).GetModuleByID(moduleId)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	_, err = new(service.NoteService).CreateNote(model.Note{
		ModuleId: module.ID,
		Content:  content,
	})
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Succ(c, "创建成功")
}

func (t *NoteController) Update(c *gin.Context) {

	_ = c.Request.ParseForm()

	id := c.PostForm("id")
	content := c.PostForm("content")
	if content == "" {
		t.Err(c, "内容不能为空")
		return
	}

	note, err := new(service.NoteService).GetNoteByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	note.Content = content
	rows := new(service.NoteService).UpdateNote(note)
	if rows == 0 {
		t.Err(c, "更新异常")
		return
	}

	t.Succ(c, "更新成功")
}

func (t *NoteController) Delete(c *gin.Context) {

	_ = c.Request.ParseForm()

	id := c.PostForm("id")
	noteService := new(service.NoteService)
	note, err := noteService.GetNoteByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	noteService.DeleteNote(note)

	t.Succ(c, "删除成功")
}