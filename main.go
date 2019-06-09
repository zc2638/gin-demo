package main

import (
	"dc-kanban/config"
	"dc-kanban/controller"
	_ "dc-kanban/lib/cron"
	_ "dc-kanban/lib/database"
	"dc-kanban/middleware"
	"dc-kanban/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/util_server"
	"net/http"
	"time"
)

func main() {

	var g = gin.Default()

	// 注册中间件
	g.Use(middleware.Cors())

	// 加载静态文件
	g.Static("/public", "./public")
	g.GET("/", new(controller.Index).Index)

	// 注册路由
	route.RouteApi(g)
	route.RouteAdmin(g)

	//开启服务
	startServer(g)
}

func startServer(g *gin.Engine) {

	server := &http.Server{
		Addr:           ":" + config.Cfg.Port,
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// 平滑退出，先结束所有在执行的任务
	util_server.GracefulExitWeb(server)
}