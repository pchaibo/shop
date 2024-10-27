package model

type User struct {
	Id         int64  `json:"id"`
	Group_id   int64  `json:"group_id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Salt       string `json:"-"`
	Createtime int64  `json:"createtime" gorm:"autoCreateTime"`
	Jointime   int64  `json:"jointime" gorm:"autoUpdateTime"`
	//Level      int64   `json:"level" gorm:"default:1"` //0:admin 1-9 user
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
func (u User) Userupdate() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}

func Usergetusername(username string) (user User, err error) {
	var u User
	u.Username = username
	res := Db.Where(u).First(&user)
	return user, res.Error
}

func Userlist(p int, username string) (user []User, err error) {
	var u []User
	page := p*10 - 10
	data := Db
	if username != "" {
		data = data.Where("username LIKE ?", "%"+username+"%")
	}
	res := data.Order("id desc").Limit(10).Offset(page).Find(&u)
	return u, res.Error
}
func Usercount() (num int64) {
	Db.Table("fa_user").Count(&num)
	return num
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
