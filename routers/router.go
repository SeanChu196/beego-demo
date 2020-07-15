package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/signin", &controllers.SigninController{})
	beego.Router("/profile", &controllers.ProfileController{})

	beego.InsertFilter("/signin", beego.BeforeRouter, controllers.GoSignUp)
}
