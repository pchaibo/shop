package admin

import (
	"fmt"
	"os"
	"shop/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Base
}

func (t LoginController) Test(c *gin.Context) {
	es, _ := os.Getwd()
	t.MakeContext(c)
	t.AjaxRun("test ok")
	fmt.Println(es)
}

//登录
func (u LoginController) Login(c *gin.Context) {
	u.MakeContext(c) //设置上下文
	c.ShouldBindJSON(&u)
	username := strings.TrimSpace(TrimHtml(u.Username))
	pass := strings.TrimSpace(TrimHtml(u.Password))
	if username == "" || pass == "" {
		u.Base.AjaxError("用户密码不能为空")
		return
	}
	data, err := model.Admingetname(username)
	if err != nil {
		u.Base.AjaxError("用户不成在")
		return

	}
	//处理pwd
	passmd5 := Md5([]byte(pass))
	if passmd5 != data.Password {
		u.Base.AjaxError("请输入正确的密码")
		return
	}
	//设置token
	token := Settoken(data.Id, data.Group_id, username)
	data.Token = token
	u.Base.AjaxRun(data)

}

//退出
func (u LoginController) Logout(c *gin.Context) {
	u.Base.MakeContext(c)
	u.AjaxRun("退出成功!")

}

//取用户信息
func (user LoginController) Info(c *gin.Context) {
	user.Base.MakeContext(c)
	token := c.Query("token")
	//fmt.Println(token)
	if token == "" {
		user.Base.AjaxError("token not noll")
		return
	}

	ken, claims, err := ParseToken(token)
	//时间超过
	if err != nil || !ken.Valid {
		user.Base.AjaxError("时间超过")
		return
	}
	//ip := Getserverip()
	//fmt.Println(claims)
	res := make(map[string]interface{})
	var avatar string
	var ros []string
	if len(claims.Username) < 1 {
		user.Base.AjaxError("用户不成在")
		return
	} else {
		//data, err := model.Admingetname(claims.Username)
		// if len(claims.Username) < 1 {
		// 	user.Base.AjaxError("用户不成在")
		// 	return
		// }
		if claims.GroupId == 1 {
			ros = append(ros, "admin")
			res["id"] = claims.GroupId
			avatar = "1654012576448720100.jpg"
		} else {
			ros = append(ros, "editor")
			avatar = "user3-128x128.jpg"
		}

	}
	res["roles"] = ros //editor admin
	res["introduction"] = "I am a super administrator"
	res["name"] = claims.Username
	host := "http://" + c.Request.Host
	res["avatar"] = host + "/imges/" + avatar

	user.Base.AjaxRun(res)

}
