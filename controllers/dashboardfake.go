package controllers

import (
	"math/rand"
	"sharit-backend/models"

	"github.com/astaxie/beego"
)

type DashboardFakeController struct {
	beego.Controller
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func (c *DashboardFakeController) Get() {
	var data []models.Point
	for i := 0; i < 800; i++ {
		var p models.Point
		p.Lat = random(41.374047, 41.413351)
		p.Lng = random(2.119863, 2.179611)
		data = append(data, p)
	}
	c.Data["data"] = data
	c.TplName = "dashboard.tpl"
}
