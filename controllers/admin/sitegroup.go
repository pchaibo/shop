package admin

import (
	"fmt"
	"html"
	"shop/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SitegroupController struct {
	Base
}

func (u SitegroupController) Add(c *gin.Context) {
	u.MakeContext(c)
	site := new(model.Sitegroup)
	jsonstr := make(map[string]interface{})
	c.BindJSON(&jsonstr)
	//fmt.Println(jsonstr)

	site.Name = jsonstr["name"].(string)

	if len(site.Name) > 1 {
		res := site.Add()
		if res > 0 {
			u.AjaxRun("添加成功!")
			return
		} else {
			u.AjaxRun("添加失败!")
			return
		}
	}

}

func (u SitegroupController) Sitegrouplist(c *gin.Context) {
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
	data, _ := model.Sitegrouplist(p, username)
	arr := make([](map[string]interface{}), len(data))
	for k, v := range data {
		arr[k] = make(map[string]interface{}) //对切片初始化
		arr[k]["id"] = v.Id
		arr[k]["status"] = v.Status
		arr[k]["name"] = v.Name
		arr[k]["remark"] = v.Remark
		createtime := time.Unix(v.Createtime, 0).Format("2006-01-02 15:04:5")
		arr[k]["createtime"] = createtime
		Updatetime := time.Unix(v.Updatetime, 0).Format("2006-01-02 15:04:5")
		arr[k]["Updatetime"] = Updatetime

	}
	resarr := make(map[string]interface{})
	resarr["total"] = model.Sitecount()
	resarr["items"] = arr
	fmt.Println("data:", arr)

	u.AjaxRun(resarr)
}

func (u SitegroupController) Del(c *gin.Context) {
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
