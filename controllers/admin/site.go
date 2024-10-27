package admin

import (
	"fmt"
	"html"
	"net/url"
	"shop/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SiteController struct {
	Base
}

func (u SiteController) Siteadd(c *gin.Context) {
	u.MakeContext(c)
	site := new(model.Site)
	jsonstr := make(map[string]interface{})
	c.BindJSON(&jsonstr)
	//fmt.Println(jsonstr)

	site.Siteurl = jsonstr["siteurl"].(string)
	urls, err := url.Parse(site.Siteurl)
	if err != nil {
		fmt.Println("url err:", err, urls)
		u.AjaxRun("添加失败!")
	}
	site.Sitename = urls.Scheme + "://" + urls.Host

	if len(site.Siteurl) > 10 {
		res := site.Useradd()
		if res > 0 {
			u.AjaxRun("添加成功!")
			return
		} else {
			u.AjaxRun("添加失败!")
			return
		}
	}

}

func (u SiteController) Sitelist(c *gin.Context) {
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
	data, _ := model.Sitelist(p, username)
	arr := make([](map[string]interface{}), len(data))
	for k, v := range data {
		arr[k] = make(map[string]interface{}) //对切片初始化
		arr[k]["id"] = v.Id
		arr[k]["status"] = v.Status
		arr[k]["sitename"] = v.Sitename
		arr[k]["siteurl"] = v.Siteurl
		arr[k]["urls"] = v.Urls
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

func (u SiteController) Del(c *gin.Context) {
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
