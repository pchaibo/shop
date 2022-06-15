package routers

import (
	"shop/controllers/admin"

	"github.com/gin-gonic/gin"
)

func Roueradmin(router *gin.Engine) *gin.Engine {
	//admins := router.Group("/admin")
	admins := router.Group("/admin", admingroup())
	{
		admins.POST("/login", admin.LoginController{}.Login)
		admins.POST("/logout", admin.LoginController{}.Logout)

		admins.GET("/info", admin.LoginController{}.Info)
		admins.GET("/user/del", admin.UserController{}.Del)
		admins.POST("/user/add", admin.UserController{}.Useradd)
		admins.GET("/user/list", admin.UserController{}.Userlist)
		admins.GET("/test", admin.LoginController{}.Test)
	}
	return router
}

//拦截
func admingroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		urlstr := url.String()
		if urlstr != "/admin/login" && urlstr != "/admin/info" {
			token := c.GetHeader("X-Token")
			if token == "" {
				c.String(200, "token not null")
				c.AbortWithStatus(200)
				return
			}
			//过滤
			_, _, err := admin.ParseToken(token)
			if err != nil {
				c.String(200, "token is error")
				c.AbortWithStatus(200)
				return
			}

		}
		//c.Next()
	}
}
