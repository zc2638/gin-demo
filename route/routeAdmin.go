package route

import (
	"dc-kanban/controller/admin"
	"dc-kanban/middleware"
	"github.com/gin-gonic/gin"
)

func RouteAdmin(g *gin.Engine) {

	admins := g.Group("/admin")

	admins.POST("/login", new(admin.Auth).Login)   // 登陆
	admins.POST("/logout", new(admin.Auth).Logout) // 登出
	admins.GET("/show", new(admin.Auth).Show)      // 详情

	admins.GET("/moduleData", new(admin.HomeController).Index)          // 模块数据
	admins.GET("/noteData", new(admin.HomeController).Note)             // 大事记数据
	admins.GET("/industryData", new(admin.HomeController).IndustryData) // 工业互联网数据接口
	admins.GET("/operateData", new(admin.HomeController).OperateData)   // IT运营数据接口
	admins.GET("/ioData", new(admin.HomeController).IOData)             // IT运营IO集合数据接口

	admins.Use(middleware.AdminAuth)
	{
		var adminController = new(admin.Admin)
		admins.GET("/admin/list", adminController.List)      // 用户列表
		admins.POST("/admin/create", middleware.AdminPermission, adminController.Create) // 用户创建
		admins.POST("/admin/update", middleware.AdminPermission, adminController.Update) // 用户更新
		admins.POST("/admin/delete", middleware.AdminPermission, adminController.Delete) // 用户删除

		var moduleController = new(admin.ModuleController)
		admins.GET("/module/list", moduleController.List)      // 模块列表
		admins.POST("/module/update", moduleController.Update) // 模块更新

		var categoryController = new(admin.CategoryController)
		admins.GET("/category/list", categoryController.List)      // 模块分类列表
		admins.GET("/category/show", categoryController.Show)      // 模块分类详情
		admins.POST("/category/create", categoryController.Create) // 模块分类创建
		admins.POST("/category/update", categoryController.Update) // 模块分类更新
		admins.POST("/category/delete", categoryController.Delete) // 模块分类删除

		var categoryItemController = new(admin.CategoryItemController)
		admins.GET("/categoryItem/list", categoryItemController.List)      // 项目列表
		admins.POST("/categoryItem/create", categoryItemController.Create) // 项目创建
		admins.POST("/categoryItem/update", categoryItemController.Update) // 项目更新
		admins.POST("/categoryItem/delete", categoryItemController.Delete) // 项目删除

		var noteController = new(admin.NoteController)
		admins.GET("/note/list", noteController.List)      // 大事记列表
		admins.POST("/note/create", noteController.Create) // 大事记创建
		admins.POST("/note/update", noteController.Update) // 大事记更新

		var configController = new(admin.ConfigController)
		admins.GET("/config/mail", configController.GetMailConfig)           // 获取邮箱配置
		admins.POST("/config/mailUpdate", configController.UpdateMailConfig) // 更新邮箱配置
	}

}
