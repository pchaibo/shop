package main

import (
	"fmt"
	"os/exec"
	"shop/controllers/common"
	"shop/routers"

	"gopkg.in/ini.v1"
)

func main() {
	//fmt.Println("ok")
	res, _ := exec.Command("go", "version").Output()
	fmt.Println(string(res))
	Start()
}

func Start() {
	conf, err := ini.Load("./conf/app.conf")
	if err != nil {
		fmt.Print(err)
	}
	ip := conf.Section("server").Key("ip").String()
	runmode := conf.Section("server").Key("runmode").String()
	s := routers.Routerinit(runmode)
	s.Run(ip)
}

//测试
func Commtest() {
	common.Getcpu()
	//common.Ffm()
	//common.PostEmail()
}
