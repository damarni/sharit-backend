package controllers

import (
	"encoding/json"
	"fmt"
	"sharit-backend/models"
)

// UserController does everything related to steam login
type UserController struct {
	BaseController
}

// Login user
func (c *UserController) Login() {

}

// PutItem register
func (c *UserController) PutItem() {
	//reb un item
	//mirar user token etc
	//afegir item
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

//EditProfile : only can update email and password
func (c *UserController) EditProfile() {

	mail := c.GetString("mail")
	myToken := c.GetString("token")
	id, err := DecodeToken(myToken)
	if err != nil {
		c.Data["json"] = "error token id"
		c.ServeJSON()
	}
	var u models.User
	u.IDuser = id
	u.Email = mail
	err = u.UpdateUser()
	if err != nil {
		fmt.Println("error al fer update")
	} else {
		fmt.Println("update ok")

	}
	// c.ServeJSON()
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
