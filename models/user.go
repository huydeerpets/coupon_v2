package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
)


//User type for authentication 
type User struct {
	ID            int64
	Email         string    `orm:"size(64);unique" json:"email" valid:"Required;Email"`
	Username	  string 	`orm:"size(64);unique" json:"username" valid:"Required;AlphaNumeric"`
	Password      string    `orm:"size(32)" json:"password" valid:"Required;MinSize(6)"`
	Role		  int 		`orm:"size(8);default(0)" json:"-"`	//normal, admin, moderate, view
	Lastlogintime time.Time `orm:"type(datetime);null" json:"-"`
	Created       time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated       time.Time `orm:"auto_now;type(datetime)" json:"-"`
}


//Valid checking valid in json body
// func (u *User) Valid(v *validation.Validation) {
// 	if u.Password != u.Repassword {
// 		v.SetError("Repassword", "Does not matched password, repassword")
// 	}
// }


//Insert create new user
func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}


//Read query user table by fields
func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}


//ReadOrCreate if read not found, then create new row
func (u *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(u, field, fields...)
}


//Update update user when modify something
func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}


//Delete delete user with conditition
func (u *User) Delete(field ...string) error {
	if _, err := orm.NewOrm().Delete(u, field...); err != nil {
		return err
	}
	return nil
}


//Users seter query when query
func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}


//init user model
func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(User))
}
