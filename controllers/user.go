package controllers

import "sharit-backend/models"

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

}

// Get get an user
func (c *UserController) Get() {

	id := c.GetString("id")
	if id == "" {
		return
	}
}
