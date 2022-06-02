package admin

import (
	"fmt"
	"net/http"
	"shop/model"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//登录
func (u LoginController) Login(c *gin.Context) {
	//username := c.PostForm("username")
	//pass := c.PostForm("password")
	c.ShouldBindJSON(&u)
	//fmt.Println(u)
	username := u.Username
	pass := u.Password
	if username == "" || pass == "" {
		out := AjaxMsg(MSG_ERR, "用户密码不能为空", "")
		c.JSON(http.StatusOK, out)
		return
	}
	data, err := model.Usergetusername(username)
	if err != nil {
		out := AjaxMsg(MSG_ERR, "data null", err)
		c.JSON(http.StatusOK, out)
		return
	}
	//处理pwd
	passmd5 := Md5([]byte(pass))
	if passmd5 != data.Password {
		out := AjaxMsg(MSG_ERR, "请输入正确的密码", err)
		c.JSON(http.StatusOK, out)
		return
	}

	//设置token
	token := Settoken(data.Id, username)
	data.Token = token
	out := AjaxMsg(MSG_OK, "ok", data)
	c.JSON(http.StatusOK, out)
}

//退出
func (u LoginController) Logout(c *gin.Context) {
	out := AjaxMsg(MSG_OK, "ok", "")
	c.JSON(http.StatusOK, out)
}

//取用户信息
func (user LoginController) Info(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)
	if token == "" {
		out := AjaxMsg(MSG_ERR, "token", "")
		c.JSON(http.StatusOK, out)
		return
	}

	ken, claims, err := ParseToken(token)
	//时间超过
	if err != nil || !ken.Valid {

		out := AjaxMsg(MSG_ERR, "时间超过", err)
		c.JSON(http.StatusOK, out)
		return
	}
	//ip := Getserverip()
	fmt.Println(claims)
	res := make(map[string]interface{})
	res["roles"] = "[admin]"
	res["introduction"] = "I am a super administrator"
	res["name"] = "holl,admin"
	host := c.Request.Host
	res["avatar"] = "http://" + host + "/static/imges/1654012576448720100.jpg"

	out := AjaxMsg(MSG_OK, "ok", res)
	c.JSON(http.StatusOK, out)

}
