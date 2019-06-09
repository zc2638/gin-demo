package cron

import (
	"dc-kanban/lib/logger"
	"dc-kanban/lib/mails"
	"dc-kanban/service"
	"fmt"
	"github.com/robfig/cron"
)

var cn = cron.New()

type mailJob struct {
	Server mails.MailServer
}

func(j mailJob) Run() {

	if err := j.Server.Send(nil); err != nil {
		logger.Info("mail Error:", err.Error())
	}
}

func init() {

	Start()
}

func Start() {

	var codes = []string{"mail_from", "mail_pwd", "mail_subject", "mail_body"}
	configList := new(service.ConfigService).GetListByCodes(codes)
	var mailConfig = make(map[string]string)
	for _, v := range configList {
		// 如果邮件配置未配全则不执行邮件计划任务
		if v.Value == "" {
			return
		}
		mailConfig[v.Code] = v.Value
	}

	moduleList := new(service.ModuleService).GetAll()
	for _, v := range moduleList {
		if v.Fixed == "" || v.Mailto == "" {
			continue
		}
		err := cn.AddJob(v.Fixed, mailJob{
			Server: mails.MailServer{
				From:    mailConfig["mail_from"],
				Pwd:     mailConfig["mail_pwd"],
				Subject: mailConfig["mail_subject"],
				Body:    mailConfig["mail_body"],
				To:      v.Mailto,
			},
		})
		if err != nil {
			fmt.Println("Cron Error:", err.Error())
		}
	}

	cn.Start()
}

func Stop() {
	cn.Stop()
}

func Restart() {
	Stop()
	Start()
}