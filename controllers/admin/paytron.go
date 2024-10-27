package admin

import (
	"fmt"
	"html"
	"shop/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PaytronController struct {
	Base
}

func (u PaytronController) Paytronlist(c *gin.Context) {
	//data, _ := model.Tronlist(1, "")
	//c.JSON(200, data)
	//
	u.MakeContext(c)
	page := c.DefaultQuery("page", "0")
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	username := c.DefaultQuery("username", "")
	if username != "" {
		username = SqlKey(username)
		username = html.EscapeString(username)
	}
	data, _ := model.Tronlist(p, username)
	arr := make([](map[string]interface{}), len(data))
	for k, v := range data {
		arr[k] = make(map[string]interface{}) //对切片初始化
		arr[k]["id"] = v.Id
		arr[k]["status"] = v.Status
		arr[k]["address"] = v.Address
		arr[k]["usdt"] = v.Usdt
		createtime := time.Unix(v.Createtime, 0).Format("2006-01-02 15:04:5")
		arr[k]["createtime"] = createtime
		Updatetime := time.Unix(v.Updatetime, 0).Format("2006-01-02 15:04:5")
		arr[k]["updatetime"] = Updatetime

	}
	resarr := make(map[string]interface{})
	resarr["total"] = model.Troncount()
	resarr["items"] = arr
	fmt.Println("data:", arr)

	u.AjaxRun(resarr)

}
