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
	beego.Router("/coupons/charge", &ctl.CouponsController{}, "post:ChargeCoupon")
	beego.Router("/coupons/get/:id", &ctl.CouponsController{}, "get:GetOne")
	beego.Router("/coupons/get", &ctl.CouponsController{}, "get:GetAll")
	beego.Router("/coupons/edit/:id", &ctl.CouponsController{}, "put:Put")
	beego.Router("/coupons/new", &ctl.CouponsController{}, "post:Post")
	beego.Router("/coupons/delete/:id", &ctl.CouponsController{}, "delete:Delete")
}
