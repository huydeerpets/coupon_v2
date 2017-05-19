package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"ttc/cps/lib"
	"ttc/cps/models"
	// "github.com/astaxie/beego"
)

// CouponsController operations for Coupons
type CouponsController struct {
	BaseController
}

// URLMapping ...
func (c *CouponsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//Prepare when access here
func (c *CouponsController) Prepare() {
	if c.Role >= lib.ROLE_NORMAL {
		c.Data["json"] = lib.Response{
			Error:       lib.ResponseUserFailPermission,
			Description: lib.ResponseUserFailPermission.String(),
		}
		c.ServeJSON()
		return
	}
}

// Post ...
// @Title Post
// @Description create Coupons
// @Param	body		body 	models.Coupons	true		"body for Coupons content"
// @Success 201 {int} models.Coupons
// @Failure 403 body is empty
// @router / [post]
func (c *CouponsController) Post() {
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

// GetOne ...
// @Title Get One
// @Description get Coupons by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Coupons
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CouponsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCouponsByID(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
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
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
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
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Coupon{ID: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCouponsByID(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
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
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCoupons(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
