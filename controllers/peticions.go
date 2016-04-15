package controllers

import (
	"encoding/json"
	"fmt"
	"sharit-backend/models"
)

// PeticionsController does everything related to steam login
type PeticionsController struct {
	BaseController
}

// Register register
func (c *PeticionsController) Register() {

	name := c.GetString("name")
	surname := c.GetString("surname")
	stars := "0"
	mail := c.GetString("mail")
	pass := c.GetString("pass")
	var u models.User
	u.IDuser = EncodeID64(mail, name, surname)
	u.Email = mail
	u.Pass = pass
	u.Name = name
	u.Stars = stars
	u.Token, _ = EncodeToken(u.IDuser, pass)
	u.Create()

	c.Data["json"] = "{\"Token\":" + u.Token + ", \"IDuser\":" + u.IDuser + "}"
	c.ServeJSON()
}

// RegisterDebug register
func (c *PeticionsController) RegisterDebug() {

	name := c.GetString("name")
	surname := c.GetString("surname")
	stars := "0"
	mail := c.GetString("mail")
	pass := c.GetString("pass")
	var u models.User
	u.IDuser = EncodeID64(mail, name, surname)
	u.Email = mail
	u.Pass = pass
	u.Name = name
	u.Stars = stars
	u.Create()
	c.ServeJSON()
}

//EditProfile : only can update email and password
func (c *PeticionsController) EditProfile() {

	mail := c.GetString("mail")
	myToken := c.GetString("token")
	id, err := DecodeToken(myToken)
	if err != nil {
		c.Data["json"] = "error token id"
		c.ServeJSON()
	}
	coordx, _ := c.GetInt("X")
	coordy, _ := c.GetInt("Y")
	var u models.User
	u.IDuser = id
	u.Email = mail
	u.X = coordx
	u.Y = coordy
	err = u.UpdateUser()
	if err != nil {
		fmt.Println("error al fer update")
	} else {
		fmt.Println("update ok")

	}
	// c.ServeJSON()
}

// GetAll get all the users
func (c *PeticionsController) GetAll() {
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
func (c *PeticionsController) Get() {

	id := c.GetString("id")

	u, err := models.FindUserByID(id)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// PutItem get a user
func (c *PeticionsController) PutItem() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	stars := "0"
	//image = ?
	token := c.GetString("token")
	iduser, _ := DecodeToken(token)
	var i models.Item
	i.ItemName = name
	i.Description = description
	i.Stars = stars

	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		u.PutItemModel(i)
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// PutItemDebug get a user
func (c *PeticionsController) PutItemDebug() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	stars := "0"
	//image = ?
	iduser := c.GetString("id")
	var i models.Item
	i.ItemName = name
	i.Description = description
	i.Stars = stars

	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		u.PutItemModel(i)
		c.Data["json"] = u
	}

	c.ServeJSON()

}
