package admin

import (
	"dc-kanban/controller"
	"dc-kanban/model"
	"dc-kanban/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct{ controller.Base }

func (t *CategoryController) List(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "15")
	moduleId := c.DefaultQuery("moduleId", "0")

	cateList, err := new(service.CategoryService).GetList(page, pageSize, moduleId)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	t.Data(c, cateList)
}

func (t *CategoryController) Show(c *gin.Context) {

	id := c.Query("id")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "15")

	itemList, err := new(service.CategoryItemService).GetListByCateID(id, page, pageSize)
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Data(c, itemList)
}

func (t *CategoryController) Create(c *gin.Context) {

	_ = c.Request.ParseForm()

	title := c.PostForm("title")
	if title == "" {
		t.Err(c, "分类标题不能为空")
	}
	desc := c.PostForm("desc")

	moduleId := c.PostForm("moduleId")
	module, err := new(service.ModuleService).GetModuleByID(moduleId)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	_, err = new(service.CategoryService).CreateCategory(model.ChartCategory{
		ModuleId: module.ID,
		Title: title,
		Desc: desc,
	})
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Succ(c, "创建成功")
}

func (t *CategoryController) Update(c *gin.Context) {

	_ = c.Request.ParseForm()

	title := c.PostForm("title")
	if title == "" {
		t.Err(c, "分类标题不能为空")
	}
	desc := c.PostForm("desc")

	id := c.PostForm("id")
	categoryService := new(service.CategoryService)
	cate, err := categoryService.GetCategoryByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	cate.Title = title
	cate.Desc = desc
	new(service.CategoryService).UpdateCategory(cate)
	t.Succ(c, "更新成功")
}

func (t *CategoryController) Delete(c *gin.Context) {

	_ = c.Request.ParseForm()

	id := c.PostForm("id")
	categoryService := new(service.CategoryService)
	cate, err := categoryService.GetCategoryByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	categoryService.DeleteCategory(cate)
	t.Succ(c, "删除成功")
}