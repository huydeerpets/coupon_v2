package controllers

import (

	"encoding/json"
	"ttc/cps/models"
	"ttc/cps/lib"
	"strings"
	"github.com/astaxie/beego"
	)


//UsersController user controller base on BaseController
type UsersController struct {
	beego.Controller
}


//CreateUser create new user
func (c *UsersController) CreateUser() {
	u := &models.User{}
	//TODO json parser to user models
	err := json.Unmarshal(c.Ctx.Input.RequestBody, u)
	if err != nil {
		return
	}
	
	if u.Username == "" && u.Email != "" {
		u.Username = strings.Split(u.Email, "@")[0]
	}
	beego.Info("%+v", u)
	if err := models.IsValid(u); err != nil {
		return
	}
	beego.Info("valid true")
	id, err := lib.CreateUser(u)
	if err != nil || id < 1 {
		return
	}

}
