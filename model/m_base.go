package model

import (
	"fmt"
	//"time"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	//取配置文件
	conf, err := ini.Load("./conf/app.conf")
	if err != nil {
		fmt.Print(err)
	}
	host := conf.Section("mysql").Key("db.host").String()    //取
	dbname := conf.Section("mysql").Key("db.name").String()  //取
	user := conf.Section("mysql").Key("db.user").String()    //取
	pwd := conf.Section("mysql").Key("db.password").String() //取

	str := "%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(str, user, pwd, host, dbname)
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "fa_", //表前缀
			SingularTable: true,  //禁用复表
		},
	})
	if error != nil {
		fmt.Println(error)
	}

	Db = db

}

var Db *gorm.DB
