package controllers

import (
	"encoding/json"
	"sharit-backend/models"
)

// ItemController does everything related to steam login
type ItemController struct {
	BaseController
}

// Put register
func (c *ItemController) Put() {

	name := c.GetString("name")
	description := c.GetString("description")

	var u models.Item
	u.ItemName = name
	u.Description = description
	u.Create()
	c.ServeJSON()
}

// GetAll get all the users
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

// GetAllRadi get an user ---- encara per implementar
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
