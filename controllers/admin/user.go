package admin

import (
	"fmt"
	"html"
	"shop/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Base
}

func (u UserController) Useradd(c *gin.Context) {
	u.MakeContext(c)
	user := new(model.User)
	jsonstr := make(map[string]interface{})
	c.BindJSON(&jsonstr)
	//fmt.Println(jsonstr)

	id := jsonstr["id"].(float64)
	uid := int64(id)

	user.Mobile = jsonstr["mobile"].(string)
	user.Username = jsonstr["username"].(string)
	user.Status = int32(jsonstr["status"].(float64))
	// if jsonstr["level"] != nil {
	// 	user.Level = int64(jsonstr["level"].(float64))
	// }
	if jsonstr["groupid"] != nil {
		user.Group_id = int64(jsonstr["groupid"].(float64))
	}

	if len(user.Mobile) != 11 {
		u.AjaxError("电话号码长度不对!")
		return
	}

	if jsonstr["password"] != nil {
		password := jsonstr["password"].(string)
		user.Password = Md5([]byte(password))

	}
	//更改
	if uid > 0 {
		user.Id = uid
		res := user.Userupdate()
		fmt.Println(res)
		if res > 0 {
			u.AjaxRun("更新成功!")
			return
		} else {
			u.AjaxError("更新失败")
			return
		}

		//添加
	} else {
		res := user.Useradd()
		if res > 0 {
			u.AjaxRun("添加成功!")
			return
		} else {
			u.AjaxRun("添加失败!")
			return
		}
	}

}

func (u UserController) Userlist(c *gin.Context) {
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
	data, _ := model.Userlist(p, username)
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
		arr[k]["Group_id"] = v.Group_id
		//arr[k]["Level"] = v.Level

	}
	resarr := make(map[string]interface{})
	resarr["total"] = model.Usercount()
	resarr["items"] = arr

	u.AjaxRun(resarr)
}

func (u UserController) Del(c *gin.Context) {
	u.MakeContext(c)
	id := c.Query("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		u.AjaxError("提交数据不对")
		return
	}
	res := model.Userdel(uid)
	fmt.Println(res)
	if res > 0 {
		u.AjaxRun("删除成功")
		return
	} else {
		u.AjaxError("删除失败")
		return
	}

}
