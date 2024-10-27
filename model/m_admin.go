package model

type Admin struct {
	Id         int64  `json:"id"`
	Username   string `json:"Username"`
	Group_id   int64
	Password   string `json:"-"`
	Salt       string `json:"-"`
	Token      string `json:"token"`
	Status     int32  `json:"status"`
	Createtime int64  `json:"createtime" gorm:"autoCreateTime"`
	Updatetime int64  `json:"updatetime" gorm:"autoUpdateTime"`
}

func (u Admin) Adminadd() int64 {
	res := Db.Create(&u)
	return res.RowsAffected
}
func (u Admin) Adminupdate() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}

func Admingetname(Adminname string) (admin Admin, err error) {
	var a Admin
	a.Username = Adminname
	res := Db.Where(a).First(&admin)
	return admin, res.Error
}

func Adminlist(p int) (u []Admin, err error) {
	page := p*10 - 10
	res := Db.Order("id desc").Limit(10).Offset(page).Find(&u)
	return u, res.Error
}
func Admincount() (num int64) {
	Db.Table("fa_Admin").Count(&num)
	return num
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
	res := Db.Model(&u).Where("id", id).Update("status", 2)
	return res.RowsAffected
}
