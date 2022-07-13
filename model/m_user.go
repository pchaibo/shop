package model

import "fmt"

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Salt       string `json:"-"`
	Createtime int64  `json:"createtime" gorm:"autoCreateTime"`
	Jointime   int64  `json:"jointime" gorm:"autoUpdateTime"`

	Updatetime int64   `json:"updatetime" gorm:"autoUpdateTime"`
	Mobile     string  `json:"mobile"`
	Token      string  `json:"token"`
	Money      float64 `json:"money"`
	Status     int32   `json:"status"`
}

func (u User) Useradd() int64 {
	res := Db.Create(&u)
	return res.RowsAffected
}
func (u User) Usserupate() int64 {
	res := Db.Model(&u).Updates(u)
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

	res := Db.Model(&u).Where("id", id).Update("status", 2)

	return res.RowsAffected
}
