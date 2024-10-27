package model

type Sitegroup struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
	Createtime int64  `json:"createtime" gorm:"autoCreateTime"`
	Updatetime int64  `json:"updatetime" gorm:"autoUpdateTime"`
	Addtime    int64  `json:"addtime" gorm:"autoUpdateTime"`
	Status     int32  `json:"status"`
}

func (u Sitegroup) Add() int64 {
	res := Db.Create(&u)
	return res.RowsAffected
}
func (u Sitegroup) Update() int64 {
	res := Db.Model(&u).Updates(u)
	return res.RowsAffected
}

func Sitegroupgetusername(sitename string) (user Sitegroup, err error) {
	var u Sitegroup
	u.Name = sitename
	res := Db.Where(u).First(&user)
	return user, res.Error
}

func Sitegrouplist(p int, username string) (user []Sitegroup, err error) {
	var u []Sitegroup
	page := p*10 - 10
	data := Db
	if username != "" {
		data = data.Where("name LIKE ?", "%"+username+"%")
	}
	res := data.Order("id desc").Limit(10).Offset(page).Find(&u)
	return u, res.Error
}
func Sitegroupcount() (num int64) {
	Db.Table("fa_sitegroup").Count(&num)
	return num
}

//删除
func Sitegroupdel(id int64) int64 {
	var u Sitegroup
	u.Id = id
	res := Db.Delete(&u)
	return res.RowsAffected
}

//停用
func Sitegrouptop(id int64) int64 {
	var u Sitegroup
	res := Db.Model(&u).Where("id", id).Update("status", 2)
	return res.RowsAffected
}
