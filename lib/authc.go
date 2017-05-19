package lib

import (
	"errors"
	"time"
	"ttc/cps/utils"
	"ttc/cps/models"
	"github.com/astaxie/beego"
)

/*
 Get authenticated user and update logintime
*/
func Authenticate(email string, password string) (user *models.User, err error) {
	msg := "invalid email or password."
	user = &models.User{Email: email}

	if err := user.Read("Email"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.ID < 1 {
		// No user
		return user, errors.New(msg)
	} else if user.Password != utils.StrTo(password).Md5() {
		// No matched password
		return user, errors.New(msg)
	} else {
		user.Lastlogintime = time.Now()
		user.Update("Lastlogintime")
		return user, nil
	}
}


//GenerateCode when need create coupons manager
func GenerateCode(length int) string {
	for {
		code := utils.GetRandomString(length)
		beego.Info(code)
		cp := models.Coupon{
			Code: code,
		}
		if models.Coupons().Filter("code", cp.Code).Exist() {
			continue
		}
		return code
	}
}
