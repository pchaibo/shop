package model

type Tron struct {
	Id         int64   `json:"id"`
	Address    string  `json:"address"`
	Usdt       float64 `json:"usdt"`
	Status     int32   `json:"status"`
	Createtime int64   `json:"createtime" gorm:"autoCreateTime"`
	Updatetime int64   `json:"updatetime" gorm:"autoUpdateTime"`
}

func Tronlist(p int, username string) (t []Tron, err error) {
	page := p*10 - 10
	data := Db
	if username != "" {
		data = data.Where("address LIKE ?", "%"+username+"%")
	}
	res := data.Order("id desc").Limit(10).Offset(page).Find(&t)
	return t, res.Error
}
func Troncount() (num int64) {
	Db.Table("fa_tron").Count(&num)
	return num
}

func (u Tron) Add() int64 {
	res := Db.Create(&u)
	return res.RowsAffected
}
