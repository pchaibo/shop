package model

import (
	"fmt"
	//"time"
)

type Admin struct {
	Id         int64   `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"-"`
	Salt       string  `json:"-"`
	Mobile     string  `json:"mobile"`
	Token      string  `json:"token"`
	Money      float64 `json:"money"`
	Status     int32   `json:"status"`
	Createtime int64   `json:"createtime" gorm:"autoCreateTime"`
	Updatetime int64   `json:"updatetime" gorm:"autoUpdateTime"`
}

func (u Admin) Useradd() int64 {

	res := Db.Create(&u)
	return res.RowsAffected
}
func (u Admin) Usserupate() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}

func Admingetusername(username string) (user Admin, err error) {
	var u Admin
	u.Username = username
	res := Db.Where(u).First(&user)

	fmt.Println(user)
	return user, res.Error
}

func Adminlist() (user []Admin, err error) {
	var u []Admin
	res := Db.Limit(10).Offset(0).Find(&u)
	return u, res.Error
}

//删除
func Admindel(id int64) int64 {
	var u Admin
	u.Id = id
	res := Db.Delete(&u)

	return res.RowsAffected
}

//停用
func Adminstop(id int64) int64 {
	var u Admin

	res := Db.Model(&u).Where("id", id).Update("status", 0)

	return res.RowsAffected
}
