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

func (u AdminController) Adminadd(c *gin.Context) {
	u.MakeContext(c)
	Admin := new(model.Admin)
	jsonstr := make(map[string]interface{})
	c.BindJSON(&jsonstr)
	//fmt.Println(jsonstr)

	id := jsonstr["id"].(float64)
	uid := int64(id)

	Admin.Username = jsonstr["username"].(string)
	Admin.Status = int32(jsonstr["status"].(float64))
	// if jsonstr["level"] != nil {
	// 	Admin.Level = int64(jsonstr["level"].(float64))
	// }
	if jsonstr["groupid"] != nil {
		Admin.Group_id = int64(jsonstr["groupid"].(float64))
	}

	if jsonstr["password"] != nil {
		password := jsonstr["password"].(string)
		Admin.Password = Md5([]byte(password))

	}
	//更改
	if uid > 0 {
		Admin.Id = uid
		res := Admin.Adminupdate()
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
		res := Admin.Adminadd()
		if res > 0 {
			u.AjaxRun("添加成功!")
			return
		} else {
			u.AjaxRun("添加失败!")
			return
		}
	}

}

func (u AdminController) Adminlist(c *gin.Context) {
	u.MakeContext(c)
	page := c.DefaultQuery("page", "0")
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	data, _ := model.Adminlist(p)
	arr := make([](map[string]interface{}), len(data))
	for k, v := range data {
		arr[k] = make(map[string]interface{})
		arr[k]["id"] = v.Id
		arr[k]["username"] = v.Username
		arr[k]["status"] = v.Status
		createtime := time.Unix(v.Createtime, 0).Format("2006-01-02 15:04:5")
		arr[k]["createtime"] = createtime
		Updatetime := time.Unix(v.Updatetime, 0).Format("2006-01-02 15:04:5")
		arr[k]["Updatetime"] = Updatetime
		arr[k]["Group_id"] = v.Group_id
		//arr[k]["Level"] = v.Level

	}
	resarr := make(map[string]interface{})
	resarr["total"] = model.Admincount()
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
	res := model.Admindel(uid)
	fmt.Println(res)
	if res > 0 {
		u.AjaxRun("删除成功")
		return
	} else {
		u.AjaxError("删除失败")
		return
	}

}
