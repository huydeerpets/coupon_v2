package routers

import (
	"github.com/astaxie/beego"
	ctl "ttc/cps/controllers"
)

func init() {
	beego.Router("/user/create", &ctl.UsersController{}, "post:CreateUser")
	beego.Router("/auth/login", &ctl.AuthController{}, "post:Login")

	//coupons manager
	beego.Router("/coupons/generatecode", &ctl.CouponsController{}, "get:GenerateCode")
}
