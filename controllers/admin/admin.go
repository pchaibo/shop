package admin

import (
	"fmt"
	"shop/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	Base
}

func (u AdminController) Useradd(c *gin.Context) {
	u.MakeContext(c)
	user := new(model.Admin)
	jsonstr := make(map[string]interface{})
	c.BindJSON(&jsonstr)
	fmt.Println(jsonstr)

	id := jsonstr["id"].(float64)
	uid := int64(id)

	user.Mobile = jsonstr["mobile"].(string)
	user.Username = jsonstr["username"].(string)
	user.Status = int32(jsonstr["status"].(float64))

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
		res := user.Add()
		if res > 0 {
			u.AjaxRun("添加成功!")
			return
		} else {
			u.AjaxError("添加失败!")
			return
		}
	}

}

func (u AdminController) Userlist(c *gin.Context) {
	u.MakeContext(c)
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

	u.AjaxRun(resarr)
}

func (u AdminController) Del(c *gin.Context) {
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
