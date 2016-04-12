package controllers

import (
	"encoding/json"
	"sharit-backend/models"
)

// UserController does everything related to steam login
type UserController struct {
	BaseController
}

// Login user
func (c *UserController) Login() {

}

// Register register
func (c *UserController) Register() {
	id := c.GetString("id")
	mail := c.GetString("mail")
	pass := c.GetString("pass")
	var u models.User
	u.IDuser = id
	u.Email = mail
	u.Pass = pass
	u.Create()
	c.ServeJSON()
}

// GetAll get all the users
func (c *UserController) GetAll() {
	users, _ := models.GetAllUsers()
	_, er := json.Marshal(users)
	if er != nil {
		//
		c.Data["json"] = "error no users"
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

// Get get a user
func (c *UserController) Get() {

	id := c.GetString("id")

	u, err := models.FindUserByID(id)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u
	}
	c.ServeJSON()

}
