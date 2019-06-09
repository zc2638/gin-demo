package admin

import (
	"dc-kanban/controller"
	"dc-kanban/lib/cron"
	"dc-kanban/model"
	"dc-kanban/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"net/http"
)

type ModuleController struct{ controller.Base }

func (t *ModuleController) List(c *gin.Context) {

	adminInfo, exist := c.Get("admin")
	if !exist {
		t.Api(c, http.StatusUnauthorized, gin.H{
			"msg": "登陆过期",
		})
		return
	}

	adminData := adminInfo.(map[string]interface{})

	var moduleList []model.Module
	if adminData["name"].(string) != "admin" {
		adminId := adminData["id"]
		admin, err := new(service.AdminService).GetAdminByID(adminId)
		if err != nil {
			t.Err(c, err.Error())
			return
		}
		var ruleMap map[string]interface{}
		if err := json.Unmarshal([]byte(admin.Rule), &ruleMap); err != nil {
			t.Err(c, "解析权限异常")
			return
		}
		moduleList = new(service.ModuleService).GetModuleByIds(ruleMap["module"])
	} else {
		moduleList = new(service.ModuleService).GetAll()
	}

	data, err := utils.ArrayStructToMap(moduleList)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	t.Data(c, data)
}

func (t *ModuleController) Update(c *gin.Context) {

	_ = c.Request.ParseForm()

	id := c.PostForm("id")
	mailto := c.PostForm("mailto")
	fixed := c.PostForm("fixed")
	if mailto == "" && fixed == "" {
		t.Err(c, "无修改内容")
		return
	}

	moduleService := new(service.ModuleService)
	module, err := moduleService.GetModuleByID(id)
	if err != nil {
		t.Err(c, err.Error())
		return
	}

	module.Mailto = mailto
	module.Fixed = fixed

	rows := moduleService.UpdateModule(module)
	if rows == 0 {
		t.Err(c, "更新异常")
		return
	}

	// 邮件自动任务重启
	cron.Restart()

	t.Succ(c, "更新成功")
}