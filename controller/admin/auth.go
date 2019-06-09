package admin

import (
	"dc-kanban/config"
	"dc-kanban/controller"
	"dc-kanban/lib/jwt"
	"dc-kanban/model"
	"dc-kanban/service"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"time"
)

type Auth struct{ controller.Base }

// 登陆
func (t *Auth) Login(c *gin.Context) {

	_ = c.Request.ParseForm()

	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" {
		t.Err(c, "请输入用户名")
		return
	}
	if password == "" {
		t.Err(c, "请输入用户名密码")
		return
	}

	admin, err := new(service.AdminService).GetAdminByName(model.Admin{
		Name:     name,
		Password: utils.MD5(password),
	})
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	var data = map[string]interface{}{
		"id":   admin.ID,
		"name": admin.Name,
		"rule": admin.Rule,
	}
	token, err := jwt.Create(data, config.JWT_SECRET_ADMIN, time.Now().Add(time.Hour * config.JWT_EXP_ADMIN).Unix())
	if err != nil {
		t.Err(c, "登陆失败")
		return
	}

	c.Request.Header.Set("token", token)
	t.Data(c, gin.H{
		"token": token,
		"name":  admin.Name,
	})
}

// 登出
func (t *Auth) Logout(c *gin.Context) {

	t.Succ(c, "登出成功")
}

// 个人详情
func (t *Auth) Show(c *gin.Context) {

	tokenStr := c.Request.Header.Get("token")
	if tokenStr == "" {
		t.Err(c, "请先登录")
		return
	}
	jwtData, err := jwt.ParseInfo(tokenStr, config.JWT_SECRET_ADMIN)
	if err != nil {
		t.Err(c, "异常登录信息1")
		return
	}
	info, ok := jwtData["info"]
	if !ok {
		t.Err(c, "异常登录信息2")
		return
	}
	id, ok := info.(map[string]interface{})["id"]
	if !ok {
		t.Err(c, "异常登录信息3")
		return
	}

	admin, err := new(service.AdminService).GetAdminByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	var data = map[string]interface{}{
		"id":   admin.ID,
		"name": admin.Name,
		"rule": admin.Rule,
	}
	token, err := jwt.Create(data, config.JWT_SECRET_ADMIN, time.Now().Add(time.Hour * config.JWT_EXP_ADMIN).Unix())
	if err != nil {
		t.Err(c, "操作失败")
		return
	}

	t.Data(c, gin.H{
		"token": token,
		"name":  admin.Name,
	})
}
