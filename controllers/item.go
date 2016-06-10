package controllers

import (
	"encoding/json"
	"sharit-backend/models"
)

type ItemController struct {
	BaseController
}

func (c *ItemController) Put() {

	name := c.GetString("name")
	description := c.GetString("description")

	var u models.Item
	u.ItemName = name
	u.Description = description
	u.Create()
	c.ServeJSON()
}

func (c *ItemController) GetAll() {
	items, _ := models.GetAllItems()
	_, er := json.Marshal(items)
	if er != nil {
		//
		c.Data["json"] = "error no items"
	} else {
		c.Data["json"] = items

	}
	c.ServeJSON()
}

func (c *ItemController) GetAllRadi() {
	x := c.GetString("x")
	y := c.GetString("y")
	if x == "" {
		return
	}
	if y == "" {
		return
	}
}
