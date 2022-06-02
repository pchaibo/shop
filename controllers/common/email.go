package common

import (
	"fmt"
	"runtime"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	Toadd string
}

func (m Mail) Email() {
	fmt.Println(m.Toadd)
	fmt.Println("toemall")
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("cpu num = ", num)
}
func PostEmail() {
	m := gomail.NewMessage()
	//m.SetHeader("From", m.FormatAddress("pchaibo@163.com", "zhang"))
	m.SetHeader("From", "zhang"+"<pchaibo@163.com>")
	m.SetHeader("To", "386378183@qq.com")
	m.SetHeader("Subject", "长江")
	html := `Hello <b>Bob</b> <a href="http://www.baidu.com/55">登录</a> and <i>Cora</i>!`
	m.SetBody("text/html", html)
	//m.Attach("./ss.docx") //附件

	d := gomail.NewDialer("smtp.163.com", 465, "pchaibo", "CJFMKMYLVZZTWKTY")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}

	//	c.String(200, "postemail"+c.ClientIP())
}
