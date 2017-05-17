package controllers

import (
	"fmt"
	"ttc/cps/lib"
	"ttc/cps/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	jwt "github.com/dgrijalva/jwt-go"
)

//Checker check token with any request
type Checker interface {
	CheckToken()
}

//BaseController define basic controller, all controller will heri in this
type BaseController struct {
	Role int
	beego.Controller
}

//var retResp util.RetObjectError

func keyFn(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(beego.AppConfig.String("token_secret")), nil

}

//CheckToken check token
func (b *BaseController) CheckToken() {
	email, err := lib.ParseToken(b.Ctx.Request.Header.Get("Authorization"))
	if err != nil {
		beego.Info("token inValid", err)
		b.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		b.ServeJSON()
	}
	beego.Info(email)
	user := &models.User{Email: email}
	beego.Info(user)
	if err := user.Read("Email"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			b.Data["json"] = lib.Response{
				Error:       lib.ResponseTokenFail,
				Description: lib.ResponseTokenFail.String(),
			}
			b.ServeJSON()
			return
		}

	}
	if user.ID < 1 {
		// No user
		b.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		b.ServeJSON()
		return
	}
	b.Role = user.Role
	return
}

//Prepare before access into this controller
func (b *BaseController) Prepare() {
	b.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	b.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, PATCH, DELETE")
	b.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type")
	if app, ok := b.AppController.(Checker); ok {
		app.CheckToken()
	}
}

//WriteResponse response json
func (b *BaseController) WriteResponse(value interface{}) {
	empty := make([]string, 0)
	switch v := value.(type) {
	case []orm.Params:
		if len(v) == 0 {
			b.Data["json"] = empty
		} else {
			b.Data["json"] = v
		}
	case []interface{}:
		if len(v) == 0 {
			b.Data["json"] = empty
		} else {
			b.Data["json"] = v
		}
	default:
		b.Data["json"] = v
	}
	b.ServeJSON()
	return
}
