package model

type Site struct {
	Id         int64  `json:"id"`
	Sitename   string `json:"sitename"`
	Siteurl    string `json:"siteurl"`
	Urls       string `json:"urls"`
	Createtime int64  `json:"createtime" gorm:"autoCreateTime"`
	Updatetime int64  `json:"updatetime" gorm:"autoUpdateTime"`
	Status     int32  `json:"status"`
}

func (u Site) Useradd() int64 {
	res := Db.Create(&u)
	return res.RowsAffected
}
func (u Site) Userupdate() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}

func Sitegetusername(sitename string) (user Site, err error) {
	var u Site
	u.Sitename = sitename
	res := Db.Where(u).First(&user)
	return user, res.Error
}

func Sitelist(p int, username string) (user []Site, err error) {
	var u []Site
	page := p*10 - 10
	data := Db
	if username != "" {
		data = data.Where("sitename LIKE ?", "%"+username+"%")
	}
	res := data.Order("id desc").Limit(10).Offset(page).Find(&u)
	return u, res.Error
}
func Sitecount() (num int64) {
	Db.Table("fa_site").Count(&num)
	return num
}

//删除
func Sitedel(id int64) int64 {
	var u Site
	u.Id = id
	res := Db.Delete(&u)
	return res.RowsAffected
}

//停用
func Sitestop(id int64) int64 {
	var u Site
	res := Db.Model(&u).Where("id", id).Update("status", 2)
	return res.RowsAffected
}
