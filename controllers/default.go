package controllers

import (
	"github.com/astaxie/beego"
)

//MainController main
type MainController struct {
	beego.Controller
}

// Get get
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.Data["json"] = "astaxie@gmail.com"
	c.ServeJSON()
	//c.ServeJSON()

}

//All aa
func (c *MainController) All() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.Data["json"] = "all"
	c.ServeJSON()
	//c.ServeJSON()

}
