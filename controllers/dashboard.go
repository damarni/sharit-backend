package controllers

import (
	"sharit-backend/models"

	"github.com/astaxie/beego"
)

type DashboardController struct {
	beego.Controller
}

//Get jhj
func (c *DashboardController) Get() {
	var data []models.Point
	data, _ = models.GetAllLogs()
	c.Data["data"] = data
	c.TplName = "dashboard.tpl"
}
