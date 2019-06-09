package admin

import (
	"dc-kanban/controller"
	"dc-kanban/model"
	"dc-kanban/service"
	"github.com/gin-gonic/gin"
)

type ConfigController struct{ controller.Base }

func (t *ConfigController) GetMailConfig(c *gin.Context) {

	var codes = []string{"mail_from", "mail_pwd", "mail_subject", "mail_body"}
	configList := new(service.ConfigService).GetListByCodes(codes)

	var data = make(gin.H)
	for _, v := range configList {
		data[v.Code] = v.Value
	}

	t.Data(c, data)
}

func (t *ConfigController) UpdateMailConfig(c *gin.Context) {

	_ = c.Request.ParseForm()

	mailFrom := c.PostForm("mailFrom")
	mailPwd := c.PostForm("mailPwd")
	mailSubject := c.PostForm("mailSubject")
	mailBody := c.PostForm("mailBody")

	var updates = make([]model.Config, 0)
	if mailFrom != "" {
		updates = append(updates, model.Config{Code: "mail_from", Value: mailFrom})
	}
	if mailPwd != "" {
		updates = append(updates, model.Config{Code: "mail_pwd", Value: mailPwd})
	}
	if mailSubject != "" {
		updates = append(updates, model.Config{Code: "mail_subject", Value: mailSubject})
	}
	if mailBody != "" {
		updates = append(updates, model.Config{Code: "mail_body", Value: mailBody})
	}
	count := len(updates)
	if count == 0 {
		t.Err(c, "无更新内容")
		return
	}

	new(service.ConfigService).BatchUpdate(updates)
	t.Succ(c, "更新成功")
}
