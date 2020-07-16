package models

import (
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var (
	ORM             orm.Ormer
	phoneRegexp     *regexp.Regexp
	emailRegexp     *regexp.Regexp
	upperCaseRegexp *regexp.Regexp
	lowerCaseRegexp *regexp.Regexp
	numRegexp       *regexp.Regexp
)

type User struct {
	Id  int    `orm:"pk;auto"`
	Usr string `orm:"unique;size(32)"`
	Pwd string `orm:"size(16)"`
}

func init() {
	//初始化持久数据库
	orm.RegisterModel(new(User))
	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	getCfg := beego.AppConfig.String
	configStr := "user=" + getCfg("pgusr") + " password=" + getCfg("pgpassword") + " dbname=" + getCfg("pgdbname") + " host=" + getCfg("pghost") + " port=" + getCfg("pgport") + " sslmode=" + getCfg("pgsslmode")
	//orm.RegisterDataBase("default", "postgres", "user=postgres password=asdf dbname=testUsr host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterDataBase("default", "postgres", configStr)
	// 自动建表
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
	ORM = orm.NewOrm()
	ORM.Using("default")
	//初始化正则式
	emailRegexp = regexp.MustCompile(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	phoneRegexp = regexp.MustCompile(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)
	upperCaseRegexp = regexp.MustCompile(`^(.*)[A-Z]+(.)*$`)
	lowerCaseRegexp = regexp.MustCompile(`^(.*)[a-z]+().*$`)
	numRegexp = regexp.MustCompile(`^(.)*[0-9]+(.)*$`)
}

// InsertUser 将用户信息插入postgres数据库
func InsertUser(usr, pwd string) (int64, error) {
	var user User
	user.Usr = usr
	user.Pwd = pwd
	//fmt.Println(ORM)
	return ORM.Insert(&user)
}

// ReadUser 从数据库读取用户信息
func ReadUser(usr *User, clo string) error {
	return ORM.Read(usr, clo)
}

//SignUp 用户注册,保存用户信息
func SignUp(usr, pwd string) bool {
	user := User{Usr: usr}
	err := ReadUser(&user, "Usr")
	if err == orm.ErrNoRows {
		//查询不到用户，可以注册
		id, er := InsertUser(usr, pwd)
		if er != nil {
			return false
		} else {
			fmt.Println("=||= Id: ", id, " Inserted!")
			return true
		}
	} else {
		fmt.Println("=||= Warning! Duplicate User ", usr)
		return false
	}
}

//GetPwd 获取用户密码
func GetPwd(usr string) (pwd string, OK bool) {
	user := User{Usr: usr}
	err := ReadUser(&user, "Usr")
	if err != orm.ErrNoRows && err != orm.ErrMissPK {
		//查询成功
		return user.Pwd, true
	} else {
		fmt.Println("=||= Query ", usr, " Failed!")
		fmt.Println(err)
		return "", false
	}
}

//VerifyPwd 验证用户名和密码
func VerifyPwd(usr, pwd string) bool {
	_pwd, OK := GetPwd(usr)
	fmt.Println("=||= Query result: ", _pwd, "  ", OK)
	if !OK || _pwd != pwd {
		return false
	}
	return true
}

//IsValidPhone 验证手机号格式
func IsValidPhone(usr string) bool {
	return phoneRegexp.MatchString(usr)
}

//IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

//IsValidPwd 验证密码格式
func IsValidPwd(pwd string) bool {
	if len(pwd) >= 8 && len(pwd) <= 16 {
		if upperCaseRegexp.MatchString(pwd) && lowerCaseRegexp.MatchString(pwd) && numRegexp.MatchString(pwd) {
			return true
		}
	}
	return false
}

//IsValidUsr 验证用户名格式
func IsValidUsr(usr string) bool {
	return IsValidEmail(usr) || IsValidPhone(usr)
}
