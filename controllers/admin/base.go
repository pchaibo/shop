package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Base struct {
	Context *gin.Context
}

//设置上下文
func (b *Base) MakeContext(c *gin.Context) {
	b.Context = c

}

func (b *Base) AjaxRun(data interface{}) {

	arr := make(map[string]interface{})
	arr["code"] = MSG_OK
	arr["message"] = "ok"
	arr["data"] = data
	b.Context.JSON(200, arr)

}

func (b *Base) AjaxError(message string) {
	arr := make(map[string]interface{})
	arr["code"] = MSG_ERR
	arr["message"] = message
	b.Context.JSON(200, arr)

}

func Test(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)
}
