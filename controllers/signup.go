package controllers

import (
	"fmt"

	"../models"
	"github.com/astaxie/beego"
)

type SignupController struct {
	beego.Controller
}

func (this *SignupController) Get() {
	this.TplName = "signup.html"
	this.Render()
	fmt.Println("=||=signup get")
}
func (this *SignupController) Post() {
	usr := this.Input().Get("usr")
	pwd := this.Input().Get("pwd")
	fmt.Println("=||=usr:", usr, "\tpwd:", pwd)
	//检查用户名和密码，添加用户信息
	if models.IsValidUsr(usr) && models.IsValidPwd(pwd) {
		if models.SignUp(usr, pwd) {
			fmt.Println("Signup Successfully")
			this.Redirect("/signin", 302)
		} else {
			fmt.Println("Store Failed!")
			this.Ctx.WriteString(`<script>alert("Signup Failed!");</script>`)
			this.TplName = "signup.html"
			this.Render()
		}
		return
	}
	//格式不正确
	fmt.Println("Wrong Format!")
	this.Ctx.WriteString(`<script>alert("Wrong Format!");</script>`)
	this.TplName = "signup.html"
	this.Render()
	return
}
