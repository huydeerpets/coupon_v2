package controllers

import (
	"github.com/astaxie/beego"
	"ttc/cps/models"
	"encoding/json"
	"ttc/cps/lib"
)


//AuthController support create user, generate Token
type AuthController struct {
	beego.Controller
}


//Login when user login
func (c *AuthController) Login() {
	u := &models.User{}
	//TODO json parser to user models
	err := json.Unmarshal(c.Ctx.Input.RequestBody, u)
	if err != nil {
		c.Data["json"] = lib.Response {
			Error: lib.ResponseJSONParseFail,
			Description: lib.ResponseJSONParseFail.String(),
		}
		c.ServeJSON()
		return
	}
	beego.Info(u)
	user, err := lib.Authenticate(u.Email, u.Password)
	if err != nil {
		c.Data["json"] = lib.Response {
			Error: lib.ResponseWrongUser,
			Description: lib.ResponseWrongUser.String(),
		}
		c.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	//generate Token
	token, err := lib.GenerateToken(user.Email)
	if err != nil {
		beego.Error(err)
	}
	data["token"] = token
	c.Data["json"] = lib.Response {
			Error: lib.ResponseOK,
			Description: lib.ResponseOK.String(),
			Data: data,
		}
	c.ServeJSON()
}