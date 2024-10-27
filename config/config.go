package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

var Tron Tronconf
var Redisconf redisconf

type Tronconf struct {
	Url string
	Key string
}
type redisconf struct {
	Ip      string
	Db      int
	Lasting string
}

var TronaddessX []Tronaddessconf

type Tronaddessconf struct {
	Add string
	Key string
}

var Ctx = context.Background()
var Rdb *redis.Client

func init() {
	conf, err := ini.Load("./conf/app.conf")
	if err != nil {
		fmt.Print(err)
	}

	Tron.Url = conf.Section("tron").Key("tronurl").String()
	Tron.Key = conf.Section("tron").Key("tronKey").String()
	Redisconf.Ip = conf.Section("redis").Key("ip").String()
	Redisconf.Lasting = conf.Section("redis").Key("lasting").String()
	Db, rediserr := conf.Section("redis").Key("db").Int()
	if rediserr != nil {
		fmt.Print("redis connet error:", rediserr)
	}
	Redisconf.Db = Db

	Resdata.Code = 2000
	Resdata.Message = "ok"
	//redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Redisconf.Ip,
		Password: "",           // no password set
		DB:       Redisconf.Db, // use default DB
	})

}

var Resdata Resdataconf

type Resdataconf struct {
	Status  int
	Code    int `json:"code"`
	Message string
	Date    interface{} `json:"date"`
}
