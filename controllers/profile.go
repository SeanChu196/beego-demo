package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

func (this *ProfileController) ProfileGet() {
	// 判断用户是否登录
	cSsid := this.Ctx.GetCookie("beegosessionID")
	fmt.Println("=||= cSSid: ", cSsid)
	usr := this.GetSession(cSsid)
	fmt.Println("=||= usr in session: ", usr)
	if usr == nil {
		fmt.Println("=||= No Session!")
		this.Redirect("/signin", 302)
		return
	}
	this.SetSession(cSsid, usr)
	fmt.Println("=||= Session: ", usr)
	this.Data["usr"] = usr
	this.TplName = "profile.html"
	this.Render()
	fmt.Println("=||= profile get")
}
