package common

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Up struct {
	Url   string
	Start int
}

func Update(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file.Header.Get("Content-Type"))

	uninx := time.Now().UnixNano()
	dst := "./static/imges/" + strconv.FormatInt(uninx, 10) + path.Ext(file.Filename)

	c.SaveUploadedFile(file, dst)

	mx := file.Size / 1024 / 1024
	log.Printf("%d MB \n", mx)
	log.Println(dst)

	//c.String(200, fmt.Sprintf("目录: %s ", dst))
	var u Up
	u.Url = dst
	out := make(map[string]interface{})
	out["code"] = 100
	out["msg"] = "msg"
	out["count"] = 20
	out["data"] = u

	c.JSON(http.StatusOK, out)

}
