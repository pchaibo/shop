package model

import (
	"fmt"
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

func (u Admin) Add() int64 {

	res := Db.Create(&u)
	return res.RowsAffected
}
func (u Admin) Upate() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}

func (u Admin) Admingetusername(username string) (user Admin, err error) {

	u.Username = username
	res := Db.Where(u).First(&user)

	fmt.Println(user)
	return user, res.Error
}

func (u Admin) Userlist() (arr []Admin, err error) {
	res := Db.Limit(10).Offset(0).Find(&arr)
	return arr, res.Error
}

//删除
func (u Admin) Del(id int64) int64 {
	u.Id = id
	res := Db.Delete(&u)

	return res.RowsAffected
}

//停用
func (u Admin) Adminstop(id int64) int64 {
	res := Db.Model(&u).Where("id", id).Update("status", 0)

	return res.RowsAffected
}
func (u Admin) Userupdate() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}
