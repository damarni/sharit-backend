package controllers

import "github.com/astaxie/beego"

// BaseController is the base controller
type BaseController struct {
	beego.Controller
}

// Prepare executes before the request
func (c *BaseController) Prepare() {

}
