package lib

import (
	"errors"
	// "fmt"
	"ttc/cps/models"
	"ttc/cps/utils"
)


//CreateUser create new user
func CreateUser(u *models.User) (int64, error) {
	var (
		err error
		msg string
	)
	if models.Users().Filter("email", u.Email).Exist() {
		msg = "was already regsitered input email address."
		return 0, errors.New(msg)
	}
	if models.Users().Filter("username", u.Username).Exist() {
		msg = "was already regsitered input username."
		return 0, errors.New(msg)
	}
	u.Password = utils.StrTo(u.Password).Md5()
	u.Role = ROLE_NORMAL
	err = u.Insert()
	if err != nil {
		return 0, err
	}
	return u.ID, err
}


//CreateCoupon create new coupon
func CreateCoupon(c *models.Coupon) (int, error) {
	var (
		err error
		msg string
	)
	if models.Coupons().Filter("code", c.Code).Exist() {
		msg = "was already regsitered this code."
		return 0, errors.New(msg)
	}
	err = c.Insert()
	if err != nil {
		return 0, err
	}
	return c.ID, err
}