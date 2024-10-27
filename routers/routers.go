package routers

import (
	"net/http"
	"shop/controllers/api"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Routerinit(runmode string) *gin.Engine {

	gin.SetMode(runmode) // debug  release 定义模式
	router := gin.Default()
	router.Use(Cors())                                //中间件
	router.SetTrustedProxies([]string{"192.168.1.2"}) //ip

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 800 << 20 // 8 MiBs
	// UI
	HTTPCacheForUI(router)
	router.StaticFile("/favicon.ico", "./web/dist/favicon.ico")
	router.StaticFS("/static", http.Dir("./web/dist/static"))
	router.StaticFS("/imges", http.Dir("./web/dist/imges"))
	router.StaticFS("/ui", http.Dir("./web/dist"))
	router.NoRoute(func(c *gin.Context) {
		log.Error().
			Str("method", c.Request.Method).
			Int("code", 404).
			Int("took", 0).
			Msg(c.Request.RequestURI)

		if strings.HasPrefix(c.Request.RequestURI, "/ui/") {
			path := strings.TrimPrefix(c.Request.RequestURI, "/ui/")
			locationPath := strings.Repeat("../", strings.Count(path, "/"))
			c.Status(http.StatusFound)
			c.Writer.Header().Set("Location", "./"+locationPath)
		} else {
			c.JSON(200, gin.H{"400": "not nill"})
		}
	})
	index := router.Group("/api")
	{
		index.POST("/Updateimage", api.Updateimage)
		index.GET("/tronaddess", api.Tronrddress)
		index.POST("/test", api.Test)
	}
	Roueradmin(router)
	//router.Run(":800")
	return router

}

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			//主要设置Access-Control-Allow-Origin
			c.Header("Access-Control-Allow-Methods", "POST, GET")
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

func HTTPCacheForUI(app *gin.Engine) {
	app.Use(func(c *gin.Context) {
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" {
			if strings.Contains(c.Request.RequestURI, "/ui/static/") {
				c.Writer.Header().Set("cache-control", "public, max-age=2592000")
				c.Writer.Header().Set("expires", time.Now().Add(time.Hour*24*30).Format(time.RFC1123))
				if strings.Contains(c.Request.RequestURI, ".js") {
					c.Writer.Header().Set("content-type", "application/javascript")
				}
				if strings.Contains(c.Request.RequestURI, ".css") {
					c.Writer.Header().Set("content-type", "text/css; charset=utf-8")
				}
			}
		}

		c.Next()
	})
}
