package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

func (this *ProfileController) Get() {
	// 判断用户是否登录
	usr := this.GetSession("Name")
	if usr == nil {
		fmt.Println("No Session!")
		this.Redirect("/signin", 302)
		return
	}
	fmt.Println("=||=Session: ", usr)
	this.Data["usr"] = usr
	this.TplName = "profile.html"
	this.Render()
	fmt.Println("profile get")
}
