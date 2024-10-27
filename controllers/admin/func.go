package admin

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/ini.v1"
)

//token
var jwtkey = []byte("55youtao.com") //设置key
//返回code
const (
	MSG_OK  = 2000
	MSG_ERR = 5000
)

func Getserverip() string {
	conf, err := ini.Load("./conf/app.conf")
	if err != nil {
		fmt.Print(err)
	}
	http := conf.Section("server").Key("http").String()

	return http
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

/**
生成密码
*/
func Password(len int, pwdO string) (pwd string, salt string) {
	salt = GetRandomString(4)
	defaultPwd := "PCHAIBO"
	if pwdO != "" {
		defaultPwd = pwdO
	}
	pwd = Md5([]byte(defaultPwd + salt))
	return pwd, salt
}

/*
HTML 过滤
*/
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile(`\\<[\\S\\s]+?\\>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile(`\\<style[\\S\\s]+?\\</style\\>`)
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile(`\\<script[\\S\\s]+?\\</script\\>`)
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile(`\\<[\\S\\s]+?\\>`)
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile(`\\s{2,}`)
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func SqlKey(str string) string {
	str = strings.Replace(str, "__", "", -1)
	return str
}

func GetRandomString(lens int) string {
	str := "123456789abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type Claims struct {
	UserId   int64
	Username string
	GroupId  int64
	jwt.StandardClaims
}

//生成token
func Settoken(uid, Groupid int64, name string) (str string) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId:   uid,
		Username: name,
		GroupId:  Groupid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1", // 签名颁发者
			//Subject:   "usertoken", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	return str
}

//解释token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
