package admin

import (
	"fmt"
	"net/http"
	"shop/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) Userlist(c *gin.Context) {
	data, _ := model.Userlist()
	arr := make([](map[string]interface{}), len(data))
	for k, v := range data {
		arr[k] = make(map[string]interface{}) //对切片初始化
		arr[k]["id"] = v.Id
		arr[k]["username"] = v.Username
		arr[k]["mobile"] = v.Mobile
		arr[k]["status"] = v.Status
		createtime := time.Unix(v.Createtime, 0).Format("2006-01-02 15:04:5")
		arr[k]["createtime"] = createtime
		Updatetime := time.Unix(v.Updatetime, 0).Format("2006-01-02 15:04:5")
		arr[k]["Updatetime"] = Updatetime
		arr[k]["money"] = v.Money

	}
	resarr := make(map[string]interface{})
	resarr["total"] = len(data)
	resarr["items"] = arr

	out := AjaxMsg(MSG_OK, "ok", resarr)
	c.JSON(http.StatusOK, out)
}

func (u UserController) Del(c *gin.Context) {
	id := c.Query("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		out := AjaxMsg(MSG_ERR, "提交数据不对!", "")
		c.JSON(http.StatusOK, out)
		return
	}
	res := model.Userdel(uid)
	fmt.Println(res)
	if res > 0 {
		out := AjaxMsg(MSG_OK, "删除成功!", id)
		c.JSON(http.StatusOK, out)
		return
	} else {
		out := AjaxMsg(MSG_ERR, "删除失败!", "")
		c.JSON(http.StatusOK, out)
		return
	}

}
