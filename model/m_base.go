package model

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Salt       string `json:"-"`
	Createtime int64  `json:"createtime,time.Time";gorm:"autoCreateTime"`
	Jointime   int64  `json:"jointime";gorm:"autoCreateTime"`

	Updatetime int64   `json:"updatetime,time.Time"`
	Mobile     int64   `json:"mobile"`
	Token      string  `json:"token"`
	Money      float64 `json:"money"`
	Status     int32   `json:"status"`
}

var Db *gorm.DB

func init() {
	//取
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

func Useradd(name string) int64 {
	var u User
	u.Username = name

	res := Db.Create(&u)
	//fmt.Println(res.RowsAffected)
	return res.RowsAffected
}

func Usergetusername(username string) (user User, err error) {
	var u User
	u.Username = username
	res := Db.Where(u).First(&user)

	fmt.Println(user)
	return user, res.Error
}

func Userlist() (user []User, err error) {
	var u []User
	res := Db.Limit(10).Offset(0).Find(&u)
	return u, res.Error
}

//删除
func Userdel(id int64) int64 {
	var u User
	u.Id = id
	res := Db.Delete(&u)

	return res.RowsAffected
}

//停用
func Userstop(id int64) int64 {
	var u User

	res := Db.Model(&u).Where("id", id).Update("status", 0)

	return res.RowsAffected
}
