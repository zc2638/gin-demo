package config

import (
	"dc-kanban/lib/logger"
	"github.com/zctod/go-tool/common/util_confd"
)

// config 配置项
type configure struct {
	Name        string `config:"default:dc-kanban;comment:项目名称"`
	Host        string `config:"default:127.0.0.1:8080;comment:项目host"`
	Port        string `config:"default:8080;comment:项目监听端口"`
	SqlHost     string `config:"default:localhost;comment:数据库地址"`
	SqlPort     string `config:"default:3306;comment:数据库端口"`
	SqlDb       string `config:"default:dc-kanban;comment:数据库名称"`
	SqlUsername string `config:"default:root;comment:数据库用户名"`
	SqlPassword string `config:"comment:数据库密码"`
}

var Cfg = &configure{}

// 初始化配置
func init() {
	if err := util_confd.InitConfig(Cfg, PATH_ENV); err != nil {
		panic(err)
	}
	logger.Info("配置初始化成功！")
}
