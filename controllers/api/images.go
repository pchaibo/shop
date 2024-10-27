package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"shop/controllers/admin"
	"strings"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	data, _ := io.ReadAll(c.Request.Body)
	//fmt.Printf("req.body=%s\n, content-type=%v\n", data, c.ContentType())

	// 这点很重要，把字节流重新放回 body 中
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

	// 获取参数
	fmt.Printf("%s \n", data)
	c.JSON(200, c.Request.Body)

}

func Updateimage(c *gin.Context) {
	str := c.PostForm("images")
	//fmt.Println(str)
	basearr := strings.Split(str, ",")
	files, err := base64.StdEncoding.DecodeString(basearr[1]) //成图片文件并把文件写入到buffer
	if err != nil {
		fmt.Println("err:", err)
	}
	jpgname := admin.GetRandomString(5)
	name := fmt.Sprintf("./static/update/%s.jpg", jpgname)
	err2 := ioutil.WriteFile(name, files, 0777)
	if err2 != nil {
		fmt.Println("err2:", err2)
	}
	base64src, pyerr := exec.Command("python", "002.py").Output()
	if pyerr != nil {
		fmt.Println("pyerr:", pyerr.Error())
	}
	fmt.Println(string(base64src))

	res := make(map[string]interface{})
	res["message"] = "ok"
	c.JSON(200, res)
}
