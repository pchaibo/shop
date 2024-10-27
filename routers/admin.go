package routers

import (
	"shop/controllers/admin"

	"github.com/gin-gonic/gin"
)

func Roueradmin(router *gin.Engine) *gin.Engine {
	//admins := router.Group("/admin")
	admins := router.Group("/v1", admingroup())
	{
		admins.POST("/login", admin.LoginController{}.Login)
		admins.POST("/logout", admin.LoginController{}.Logout)

		admins.GET("/info", admin.LoginController{}.Info)
		admins.GET("/user/del", admin.UserController{}.Del)
		admins.POST("/user/add", admin.UserController{}.Useradd)
		admins.GET("/user/list", admin.UserController{}.Userlist)
		//
		admins.GET("/admin/del", admin.AdminController{}.Del)
		admins.POST("/admin/add", admin.AdminController{}.Adminadd)
		admins.GET("/admin/list", admin.AdminController{}.Adminlist)
		admins.GET("/test", admin.LoginController{}.Test)
		//
		admins.GET("/site/list", admin.SiteController{}.Sitelist)
		admins.POST("/site/add", admin.SiteController{}.Siteadd)
		//sitegroup
		admins.GET("/sitegroup/list", admin.SitegroupController{}.Sitegrouplist)
		admins.POST("/sitegroup/add", admin.SitegroupController{}.Add)
		//Paytronlist
		admins.GET("/paytron/list", admin.PaytronController{}.Paytronlist)
	}
	return router
}

// 拦截
func admingroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		urlstr := url.String()
		if urlstr != "/v1/login" && urlstr != "/v1/info" {
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
