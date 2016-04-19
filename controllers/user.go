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

// Register register
func (c *UserController) Register() {

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
func (c *UserController) RegisterDebug() {

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
func (c *UserController) EditProfile() {

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

// PutItem get a user
func (c *UserController) PutItem() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	stars := "0"
	//image = ?
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
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

// GetItems return user items
func (c *UserController) GetItems() {
	token := c.GetString("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u.ItemsUser
	}
	c.ServeJSON()

}

// PutItemDebug get a user
func (c *UserController) PutItemDebug() {
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

// PutPeticioRadi put peticio al radi
func (c *UserController) PutPeticioRadi() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	var p models.Peticio
	p.IDuser = iduser
	p.Name = name
	p.To = ""
	p.Descripcio = description
	p.X = u.X
	p.Y = u.Y
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		p.Create()
		c.Data["json"] = p
	}
	c.ServeJSON()
}

// PutPeticioRadiDebug get a user
func (c *UserController) PutPeticioRadiDebug() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	iduser := c.GetString("iduser")

	u, err := models.FindUserByID(iduser)
	var p models.Peticio
	p.IDuser = iduser
	p.Name = name
	p.To = ""
	p.Descripcio = description
	p.X = u.X
	p.Y = u.Y
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		p.Create()
		c.Data["json"] = p
	}
	c.ServeJSON()
}

// PutPeticioUsuari get a user
func (c *UserController) PutPeticioUsuari() {
	//fer una peticio especifica a un usuari

}

// GetPeticionsRadiUser get a user
func (c *UserController) GetPeticionsRadiUser() {
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		peticions, err := models.GetPeticionsRadi(u.X, u.Y)
		if err == nil {
			c.Data["json"] = peticions
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// GetPeticionsUsuari get a user
func (c *UserController) GetPeticionsUsuari() {
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		c.Data["json"] = u.PeticionsUser
		c.ServeJSON()

	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}
