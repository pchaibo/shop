package routers

import (
	"net/http"
	"shop/controllers/common"

	"github.com/gin-gonic/gin"
)

func Routerinit(runmode string) *gin.Engine {

	gin.SetMode(runmode) // debug  release 定义模式
	router := gin.Default()
	router.Use(Cors()) //中间件

	//router := gin.New()
	router.SetTrustedProxies([]string{"192.168.1.2"}) //ip

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 800 << 20 // 8 MiBs
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/s", "./static/index.html")

	index := router.Group("/v1")
	{
		index.POST("/update", common.Update)
	}
	Roueradmin(router)

	router.NoRoute(func(c *gin.Context) {

		c.String(200, "您的请求404！")
	})

	//router.Run(":800")
	return router

}

//跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			//主要设置Access-Control-Allow-Origin
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			//停重复请求
			c.Header("Access-Control-Max-Age", "2592000")
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
