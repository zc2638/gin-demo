package admin

import (
	"dc-kanban/controller"
	"dc-kanban/model"
	"dc-kanban/service"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"strconv"
)

type CategoryItemController struct{ controller.Base }

func (t *CategoryItemController) List(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "15")
	categoryId := c.DefaultQuery("categoryId", "0")

	itemList, err := new(service.CategoryItemService).GetListByCateID(categoryId, page, pageSize)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	data, err := utils.ArrayStructToMap(itemList)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	t.Data(c, data)
}

func (t *CategoryItemController) Create(c *gin.Context) {

	_ = c.Request.ParseForm()

	categoryId := c.PostForm("categoryId")
	cate, err := new(service.CategoryService).GetCategoryByID(categoryId)
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	value := c.DefaultPostForm("value", "")

	typ := c.DefaultPostForm("type", "0")
	typN, err := strconv.Atoi(typ)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	sort := c.DefaultPostForm("sort", "1")
	sortN, err := strconv.Atoi(sort)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	_, err = new(service.CategoryItemService).CreateItem(model.ChartItem{
		CategoryId: cate.ID,
		Type:       uint(typN),
		Title:      title,
		Content:    content,
		Value:      value,
		Sort:       uint(sortN),
	})
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Succ(c, "添加成功")
}

func (t *CategoryItemController) Update(c *gin.Context) {

	_ = c.Request.ParseForm()

	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	value := c.DefaultPostForm("value", "")

	typ := c.DefaultPostForm("type", "0")
	typN, err := strconv.Atoi(typ)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	sort := c.DefaultPostForm("sort", "1")
	sortN, err := strconv.Atoi(sort)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	id := c.PostForm("id")
	itemService := new(service.CategoryItemService)
	item := itemService.GetItemByID(id)
	if item.ID == 0 {
		t.Err(c, "不存在的项目")
		return
	}

	item.Title = title
	item.Content = content
	item.Value = value
	item.Type = uint(typN)
	item.Sort = uint(sortN)
	rows := itemService.UpdateItem(item)
	if rows == 0 {
		t.Err(c, "更新异常")
		return
	}

	t.Succ(c, "更新成功")
}

func (t *CategoryItemController) Delete(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.PostForm("id")

	itemService := new(service.CategoryItemService)
	item := itemService.GetItemByID(id)
	if item.ID == 0 {
		t.Err(c, "不存在的项目")
		return
	}

	itemService.DeleteItem(item)
	t.Succ(c, "删除成功")
}