package routers

import (
	"github.com/astaxie/beego"
	"music-front/controllers"
)

func init() {
	beego.Router("/signUp", &controllers.SignUpController{}, "post:SignUp")
}
