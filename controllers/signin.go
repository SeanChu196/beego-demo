package controllers

import (
	"beego-demo/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type SigninController struct {
	beego.Controller
}

func (this *SigninController) SigninGet() {
	this.TplName = "signin.html"
	this.Render()
	fmt.Println("=||= signin get")
}
func (this *SigninController) SigninPost() {
	usr := this.Input().Get("usr")
	pwd := this.Input().Get("pwd")
	fmt.Println("=||= usr:", usr, "\tpwd:", pwd)
	//检查用户名和密码设置登录信息
	if models.IsValidUsr(usr) && models.IsValidPwd(pwd) {
		if models.VerifyPwd(usr, pwd) {
			cSsid := this.Ctx.GetCookie("beegosessionID")
			fmt.Println(cSsid)
			Ssid := this.GetSession("beegosessionID")
			fmt.Println(Ssid)
			fmt.Println("******************")
			//添加Session登录信息
			this.SetSession(cSsid, usr)
			fmt.Println("Set Session!")
			fmt.Println("=||= getsession: ", this.GetSession(cSsid))
			fmt.Println("-----------------")
			//this.Ctx.SetCookie("Usr", usr)
			this.Redirect("/profile", 302)
		} else {
			fmt.Println("=||= Wrong Password!")
			this.Ctx.WriteString(`<script>alert("Incorrect user name or password!");</script>`)
			this.TplName = "signin.html"
			this.Render()
		}
		return
	}
	//登录失败
	fmt.Println("=||= login Failed!")
	this.Ctx.WriteString(`<script>alert("Login Failed!");</script>`)
	this.TplName = "signin.html"
	this.Render()
	return
}
func (this *SigninController) SigninPut() {
	fmt.Println("=||= go to signup")
	this.Redirect("/signup", 302)
}
func GoSignUp(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = ctx.Input.Query("_method")
	}
}
