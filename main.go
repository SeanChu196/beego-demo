package main

import (
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	"beego-demo/controllers"

	_ "github.com/lib/pq"
)

func init() {
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 60
	//路由设置
	beego.Router("/signup", &controllers.SignupController{}, "get:SignupGet;post:SignupPost")
	beego.Router("/signin", &controllers.SigninController{}, "get:SigninGet;post:SigninPost;put:SigninPut")
	beego.Router("/profile", &controllers.ProfileController{}, "get:ProfileGet")
	beego.InsertFilter("/signin", beego.BeforeRouter, controllers.GoSignUp)
	fmt.Println("=||= main.init() OK")
}

func main() {
	beego.Run()
}
