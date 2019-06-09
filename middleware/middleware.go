package middleware

import (
	"dc-kanban/config"
	"dc-kanban/lib/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域支持
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, X-Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 身份认证
func AdminAuth(c *gin.Context) {

	tokenStr := c.Request.Header.Get("token")
	if jwt.CheckValid(tokenStr, config.JWT_SECRET_ADMIN) == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "登陆失败",
		})
		return
	}

	jwtData, err := jwt.ParseInfo(tokenStr, config.JWT_SECRET_ADMIN)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "异常登录信息",
		})
		return
	}

	info, ok := jwtData["info"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "异常登录信息",
		})
		return
	}
	c.Set("admin", info.(map[string]interface{}))

	c.Next()
}

func AdminPermission(c *gin.Context) {

	adminInfo ,exist := c.Get("admin")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "登陆失败",
		})
		return
	}
	adminData := adminInfo.(map[string]interface{})
	if adminData["id"].(float64) != 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "无操作权限",
		})
	}
	c.Next()
}