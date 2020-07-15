package controllers

import (
	"fmt"

	"../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type SigninController struct {
	beego.Controller
}

func (this *SigninController) Get() {
	this.TplName = "signin.html"
	this.Render()
	fmt.Println("=||=signin get")
}
func (this *SigninController) Post() {
	usr := this.Input().Get("usr")
	pwd := this.Input().Get("pwd")
	fmt.Println("=||=usr:", usr, "\tpwd:", pwd)
	//检查用户名和密码设置登录信息
	if models.IsValidUsr(usr) && models.IsValidPwd(pwd) {
		if models.VerifyPwd(usr, pwd) {
			//添加登录信息
			this.SetSession("Name", usr)
			fmt.Println("Set Session!")
			this.Redirect("/profile", 302)
		} else {
			fmt.Println("Wrong Password!")
			this.Ctx.WriteString(`<script>alert("Incorrect user name or password!");</script>`)
			this.TplName = "signin.html"
			this.Render()
		}
		return
	}
	//登录失败
	fmt.Println("login Failed!")
	this.Ctx.WriteString(`<script>alert("Login Failed!");</script>`)
	this.TplName = "signin.html"
	this.Render()
	return
	//this.Data["jsonp"] = &mystruct
	//this.ServeJSONP()
	//("错误")
}
func (this *SigninController) Put() {
	fmt.Println("=||=go to signup")
	this.Redirect("/signup", 302)
}
func GoSignUp(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = ctx.Input.Query("_method")
	}
}
