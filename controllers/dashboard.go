package controllers

import (
	"math/rand"

	"github.com/astaxie/beego"
)

type DashboardController struct {
	beego.Controller
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lon"`
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

//Get jhj
func (c *DashboardController) Get() {
	var data []Point
	for i := 0; i < 800; i++ {
		var p Point
		p.Lat = random(41.374047, 41.413351)
		p.Lng = random(2.119863, 2.179611)
		data = append(data, p)
	}
	c.Data["data"] = data
	c.TplName = "dashboard.tpl"
}
