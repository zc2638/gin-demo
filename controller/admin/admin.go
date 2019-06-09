package admin

import (
	"dc-kanban/controller"
	"dc-kanban/model"
	"dc-kanban/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"strconv"
)

type Admin struct{ controller.Base }

// 后台首页
func (t *Admin) Index(c *gin.Context) {
	t.Succ(c, "Hello World!")
}

// 管理员列表
func (t *Admin) List(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "15")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		t.Err(c, "请填写page为数值类型")
		return
	}
	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		t.Err(c, "请填写pageSize为数值类型")
		return
	}

	adminList := new(service.AdminService).GetList(pageNum, pageSizeNum)
	data, err := utils.ArrayStructToMap(adminList)
	if err != nil {
		t.Err(c, "解析出错")
		return
	}

	t.Data(c, data)
}

// 管理员添加
func (t *Admin) Create(c *gin.Context) {

	_ = c.Request.ParseForm()
	name := c.PostForm("name")
	password := c.PostForm("password")
	rule := c.PostForm("rule")

	if name == "" {
		t.Err(c, "请输入用户名称")
		return
	}
	if password == "" {
		t.Err(c, "请输入用户密码")
		return
	}

	var ruleMap map[string]interface{}
	if err := json.Unmarshal([]byte(rule), &ruleMap); err != nil {
		t.Err(c, "权限格式异常")
		return
	}

	_, err := new(service.AdminService).CreateAdmin(model.Admin{
		Name:     name,
		Password: password,
		Rule:     rule,
	})
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Succ(c, "创建成功")
}

// 管理员修改
func (t *Admin) Update(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.DefaultPostForm("id", "0")
	name := c.PostForm("name")
	password := c.PostForm("password")
	if id == "0" {
		t.Err(c, "请选择用户")
		return
	}
	if name == "" && password == "" {
		t.Err(c, "无修改内容")
		return
	}

	adminService := new(service.AdminService)
	admin, err := adminService.GetAdminByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	if name != "" {
		_, err := adminService.GetAdminByName(model.Admin{
			Name: name,
		})
		if err == nil {
			t.Err(c, "该用户名称已被占用")
			return
		}
		admin.Name = name
	}
	if password != "" {
		admin.Password = utils.MD5(password)
	}
	adminService.UpdateAdmin(admin)

	t.Succ(c, "修改成功")
}

// 管理员删除
func (t *Admin) Delete(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.PostForm("id")
	if id == "1" {
		t.Err(c, "超级用户无法删除")
		return
	}
	adminService := new(service.AdminService)
	admin, err := adminService.GetAdminByID(id)
	if err != nil {
		t.Err(c, "不存在的用户")
		return
	}

	adminService.DeleteAdmin(admin)

	t.Succ(c, "删除成功")
}
