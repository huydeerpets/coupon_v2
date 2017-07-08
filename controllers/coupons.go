package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"ttc/cps/lib"
	"ttc/cps/models"
	"ttc/cps/utils"
	log "ttc/utilities/logging"
	// "github.com/astaxie/beego"
)

// CouponsController operations for Coupons
type CouponsController struct {
	BaseController
}

const (
	activeSTATUS = 1
)

// URLMapping ...
func (c *CouponsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//ValidRole when access here
func (c *CouponsController) ValidRole() bool {
	if c.Role >= lib.ROLE_NORMAL {
		return true
	}
	return false
}

// Post ...
// @Title Post
// @Description create Coupons
// @Param	body		body 	models.Coupons	true		"body for Coupons content"
// @Success 201 {int} models.Coupons
// @Failure 403 body is empty
// @router / [post]
func (c *CouponsController) Post() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}
	var v models.Coupon
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCoupon(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//GenerateCode ...
// @Title GenerateCode
// @Description generate code
// @Success 200 {object} code
// @router /generatecode [get]
func (c *CouponsController) GenerateCode() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	code := lib.GenerateCode(6)
	data["code"] = code
	c.Data["json"] = lib.Response{
		Error:       lib.ResponseOK,
		Description: lib.ResponseOK.String(),
		Data:        data,
	}
	c.ServeJSON()

}

//ChargeCoupon ...
// @Title ChargeCoupon
// @Description generate code
// @Success 200 {object} code
// @router /charge [post]
func (c *CouponsController) ChargeCoupon() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}

	//TODO: check json code
	cp := models.Coupon{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cp)
	if err != nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseJSONParseFail,
			Description: lib.ResponseJSONParseFail.String(),
		}
		c.ServeJSON()
		return
	}
	log.Info("%+v", cp)
	//TODO: check code exist in Database
	// cpdb := models.Coupon{}
	cpdb, err := models.GetCouponsByCode(cp.Code)
	if err != nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseCodeInvalid,
			Description: lib.ResponseCodeInvalid.String(),
		}
		c.ServeJSON()
		return
	}
	log.Info("cp from db %+v", cpdb)
	//TODO: check valid count limit
	if cpdb.RedemptionLimit < cpdb.CouponRedemptionsCount && cpdb.CouponRedemptionsCount > 0 {
		log.Info("coupon da het so lan su dung")
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseCodeInvalid,
			Description: lib.ResponseCodeInvalid.String(),
		}
		c.ServeJSON()
		return
	}
	//TODO: check valid time until
	log.Info("time valid %s", cpdb.ValidFrom)

	//TODO: check valid status
	if cpdb.Status != activeSTATUS {
		log.Info("coupon da bi disable")

		c.Data["json"] = lib.Response{
			Error:       lib.ResponseCodeInvalid,
			Description: lib.ResponseCodeInvalid.String(),
		}
		c.ServeJSON()
		return
	}
	//TODO: check valid username
	if cp.Username != cpdb.Username {
		log.Info("coupon khong danh cho ban")

		c.Data["json"] = lib.Response{
			Error:       lib.ResponseCodeInvalid,
			Description: lib.ResponseCodeInvalid.String(),
		}
		c.ServeJSON()
		return
	}
	//TODO: check valid categories
	cates := strings.Split(cpdb.Categories, ",")
	if cpdb.Categories != "" && !utils.StringInSlice(cp.Categories, cates) {
		log.Info("coupon khong trong cates nay")

		c.Data["json"] = lib.Response{
			Error:       lib.ResponseCodeInvalid,
			Description: lib.ResponseCodeInvalid.String(),
		}
		c.ServeJSON()
		return
	}
	//TODO: check valid products

	products := strings.Split(cpdb.Products, ",")
	if cpdb.Products != "" && !utils.StringInSlice(cp.Products, products) {
		log.Info("coupon khong ap dung cho san pham nay")

		c.Data["json"] = lib.Response{
			Error:       lib.ResponseCodeInvalid,
			Description: lib.ResponseCodeInvalid.String(),
		}
		c.ServeJSON()
		return
	}

	//TODO: response OK
	c.Data["json"] = lib.Response{
		Error:       lib.ResponseOK,
		Description: lib.ResponseOK.String(),
	}
	cpdb.CouponRedemptionsCount++
	cpdb.Update("coupon_redemptions_count")
	// inc Coupon.CouponRedemptionsCount by 1
}

// GetOne ...
// @Title Get One
// @Description get Coupons by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Coupons
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CouponsController) GetOne() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCouponsByID(id)
	if err != nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseDatabaseNotFoundCoupon,
			Description: lib.ResponseDatabaseNotFoundCoupon.String(),
		}
	} else {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseOK,
			Description: lib.ResponseOK.String(),
			Data:        v,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Coupons
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Coupons
// @Failure 403
// @router / [get]
func (c *CouponsController) GetAll() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCoupons(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseDatabaseNotFoundCoupon,
			Description: lib.ResponseDatabaseNotFoundCoupon.String(),
		}
	} else {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseOK,
			Description: lib.ResponseOK.String(),
			Data:        l,
		}
	}
	c.ServeJSON()
}

//PrepareUpdate check conditition for update
func PrepareUpdate(code string) (v models.Coupon, err error) {
	cp := models.Coupon{Code: code}
	err = cp.Read("code")
	if err.Error() == "<QuerySeter> no row found" || err != nil {
		return v, err
	}
	//check fields
	if cp.Amount != 0 {
		v.Amount = cp.Amount
	}
	if cp.ValidFrom.String() != "0001-01-01T00:00:00Z" {
		v.ValidFrom = cp.ValidFrom
	}
	if cp.ValidUntil.String() != "0001-01-01T00:00:00Z" {
		v.ValidUntil = cp.ValidUntil
	}
	if cp.Categories != "" {
		v.Categories = cp.Categories
	}
	if cp.Description != "" {
		v.Description = cp.Description
	}
	if cp.Products != "" {
		v.Products = cp.Products
	}
	if cp.Username != "" {
		v.Username = cp.Username
	}
	if cp.RedemptionLimit != 0 {
		v.RedemptionLimit = cp.RedemptionLimit
	}
	if cp.Status != 0 {
		v.Status = cp.Status
	}
	if cp.Type != 0 {
		v.Type = cp.Type
	}
	return v, nil
}

// Put ...
// @Title Put
// @Description update the Coupons
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Coupons	true		"body for Coupons content"
// @Success 200 {object} models.Coupons
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CouponsController) Put() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}
	cp := models.Coupon{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cp)
	if err != nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseJSONParseFail,
			Description: lib.ResponseJSONParseFail.String(),
		}
		c.ServeJSON()
		return
	}
	v, err := PrepareUpdate(cp.Code)
	if err.Error() == "<QuerySeter> no row found" || err != nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseDatabaseNotFoundCoupon,
			Description: lib.ResponseDatabaseNotFoundCoupon.String(),
		}
		c.ServeJSON()
		return
	}
	v.Update("code")
	c.Data["json"] = lib.Response{
		Error:       lib.ResponseOK,
		Description: lib.ResponseOK.String(),
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Coupons
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CouponsController) Delete() {
	if !c.ValidRole() {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseTokenFail,
			Description: lib.ResponseTokenFail.String(),
		}
		c.ServeJSON()
		return
	}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCoupons(id); err == nil {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseOK,
			Description: lib.ResponseOK.String(),
		}
	} else {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseDatabaseErrorCoupon,
			Description: lib.ResponseDatabaseErrorCoupon.String(),
		}
	}
	c.ServeJSON()
}
