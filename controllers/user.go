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

}

// GetAll get all the users
func (c *UserController) GetAll() {

}

// Get get an user
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
