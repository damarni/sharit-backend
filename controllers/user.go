package controllers

import (
	"encoding/json"
	"fmt"
	"sharit-backend/models"
	"strconv"
	"time"
)

// UserController does everything related to steam login
type UserController struct {
	BaseController
}

type LoginStruct struct {
	Email string `bson:"email"`
	X     int    `bson:"x"`
	Y     int    `bson:"y"`
	Pass  string `bson:"pass"`
}

// Login user
func (c *UserController) Login() {
	mail := c.GetString("email")

	pass := c.GetString("pass")
	u, err := models.FindUserByMail(mail)
	if err == nil {
		if pass == u.Pass {
			var r reg
			r.Token = u.Token
			r.Iduser = u.IDuser
			c.Data["json"] = r
			c.ServeJSON()
		} else {
			c.Data["json"] = "wrong pass"
			c.ServeJSON()
		}

	} else {
		c.Data["json"] = "mail not registered"
		c.ServeJSON()
	}
}

type reg struct {
	Token  string `bson:"token,omitempty"`
	Iduser string `bson:"iduser,omitempty"`
}

// DeleteUser get a user
func (c *UserController) SendOptions() {

	token := c.Ctx.Input.Header("token")
	idToken, err := DecodeToken(token)

	if err == nil {

		err = models.DeleteUserByID(idToken)

		if err != nil {
			fmt.Println(err)
			c.Data["json"] = "user not found"
		} else {
			c.Data["json"] = "user deleted"
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = "token fail"
		c.ServeJSON()
	}

}

// Register register
func (c *UserController) Register() {
	var datapoint models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	name := datapoint.Name
	surname := datapoint.Surname
	stars := "0"
	mail := datapoint.Email
	pass := datapoint.Pass
	image := datapoint.Image
	fmt.Println(datapoint)
	_, err := models.FindUserByID(EncodeID64(mail, name, surname))
	if err != nil {
		var u models.User
		u.IDuser = EncodeID64(mail, name, surname)
		u.Email = mail
		u.Surname = surname
		u.Pass = pass
		u.Name = name
		u.Stars = stars
		u.Image = image
		coordx := 1
		coordy := 1
		u.X = coordx
		u.Y = coordy
		u.Token, _ = EncodeToken(u.IDuser, pass)
		u.Create()
		var r reg
		r.Token = u.Token
		r.Iduser = u.IDuser
		c.Data["json"] = r
		c.ServeJSON()
	} else {
		c.Data["json"] = "mail used"
		c.ServeJSON()
	}

}

// Options register
func (c *UserController) Options() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.ServeJSON()

}

//EditProfile : only can update email and password
func (c *UserController) EditProfile() {
	var datapoint LoginStruct
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	fmt.Println(datapoint.Email)
	mail := datapoint.Email
	token := c.Ctx.Input.Header("token")
	id, err := DecodeToken(token)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = "error token id"
		c.ServeJSON()
	}
	coordx := datapoint.X
	coordy := datapoint.Y
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

	c.Data["json"] = u
	c.ServeJSON()
}

// GetAll get all the users
func (c *UserController) GetAll() {
	fmt.Println("getall")
	users, err := models.GetAllUsers()
	fmt.Println(err)
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

	token := c.Ctx.Input.Header("token")
	idToken, err := DecodeToken(token)

	if err == nil {
		id := c.GetString("id")
		var u models.User
		if id != "" {
			u, err = models.FindUserByID(id)
		} else {
			u, err = models.FindUserByID(idToken)
		}
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = "user not found"
		} else {
			c.Data["json"] = u
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = "token fail"
		c.ServeJSON()
	}

}

// DeleteUser get a user
func (c *UserController) DeleteUser() {

	token := c.Ctx.Input.Header("token")
	idToken, err := DecodeToken(token)

	if err == nil {

		err = models.DeleteUserByID(idToken)

		if err != nil {
			fmt.Println(err)
			c.Data["json"] = "user not found"
		} else {
			c.Data["json"] = "user deleted"
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = "token fail"
		c.ServeJSON()
	}

}

type PetDel struct {
	IDPet string `bson:"idPeticio"`
}

// DeletePeticio get a user
func (c *UserController) DeletePeticio() {
	var datapoint models.Peticio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	token := c.Ctx.Input.Header("token")
	_, err := DecodeToken(token)
	idpet := datapoint.ID
	fmt.Println(datapoint)
	if err == nil {

		err = models.DeletePeticioByID(idpet)
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = "peticio not found"
		} else {
			c.Data["json"] = "peticio deleted"
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = "token fail"
		c.ServeJSON()
	}

}

// PutItem get a user
func (c *UserController) PutItem() {
	var datapoint models.Item
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	name := datapoint.ItemName
	description := datapoint.Description
	stars := "0"
	image1 := datapoint.Image1
	image2 := datapoint.Image2
	image3 := datapoint.Image3
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	var i models.Item
	stt := token + name + time.Now().String()
	i.IDuser = iduser
	i.Idd = EncodeMsg(stt)
	i.ItemName = name
	i.Description = description
	i.Stars = stars
	i.Image1 = image1
	i.Image2 = image2
	i.Image3 = image3
	i.LastSharit = time.Now()
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		fmt.Println("ok item")
		u.PutItemModel(i)
		u, _ := models.FindUserByID(iduser)
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// DeleteItem get a user
func (c *UserController) DeleteItem() {
	var datapoint models.Item
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	idItem := datapoint.Idd
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)

	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		fmt.Println("ok item")
		u.DeleteItemModel(idItem)
		u, _ := models.FindUserByID(iduser)
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// GetItems return user items
func (c *UserController) GetItems() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u.ItemsUser
	}
	c.ServeJSON()

}

// GetTransaccions return user items
func (c *UserController) GetTransaccions() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u.Transaccions
	}
	c.ServeJSON()

}

// GetItemsRadi return user items
func (c *UserController) GetItemsRadi() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	fmt.Println(iduser)
	if err == nil {
		items, err := models.GetItemsRadi(u.X, u.Y)
		if err == nil {
			c.Data["json"] = items
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// UpdateItem update user item
func (c *UserController) UpdateItem() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)

	var datapoint models.Item
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	name := datapoint.ItemName
	iditem := datapoint.Idd
	description := datapoint.Description
	image1 := datapoint.Image1
	var it models.Item
	it.Idd = iditem
	it.Description = description
	it.ItemName = name
	it.Image1 = image1
	err = u.UpdateItemModels(it)
	if err == nil {
		c.Data["json"] = it
	} else {
		c.Data["json"] = "error updating"
	}
	c.ServeJSON()

}

type GetItemStruct struct {
	IDItem string `bson:"idItem"`
	IDUser string `bson:"idUser"`
}

// GetItem return user items
func (c *UserController) GetItem() models.Item {
	token := c.Ctx.Input.Header("token")
	_, err := DecodeToken(token)
	var datapoint GetItemStruct
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	idItem := datapoint.IDItem
	idUser := datapoint.IDUser
	u, err := models.FindUserByID(idUser)
	var item models.Item
	uintID, _ := strconv.ParseUint(idItem, 10, 32)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		items := u.ItemsUser
		for _, it := range items {
			if it.ID == uintID {
				item = it
			}
		}
	}
	return item
}

// GetItemSoft return user items
func (c *UserController) GetItemSoft() {
	token := c.Ctx.Input.Header("token")
	idUser, err := DecodeToken(token)

	idItem := c.GetString("idItem")
	u, err := models.FindUserByID(idUser)
	var item models.Item
	uintID, _ := strconv.ParseUint(idItem, 10, 32)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		items := u.ItemsUser
		for _, it := range items {
			if it.ID == uintID {
				item = it
			}
		}
	}
	c.Data["json"] = item

	c.ServeJSON()
}

// PutPeticioUsuari get a user
func (c *UserController) PutPeticio() {
	//fer una peticio especifica a un usua
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)

	var datapoint models.Peticio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	name := datapoint.Name
	description := datapoint.Descripcio
	image := datapoint.Image
	u, _ := models.FindUserByID(iduser)
	var p models.Peticio
	p.IDuser = iduser
	p.UserName = u.Name
	p.UserSurname = u.Surname
	p.Image = image
	p.ID = EncodeMsg(iduser + time.Now().String())
	p.Name = name
	p.To = ""
	p.Descripcio = description
	p.X = u.X
	p.Y = u.Y
	p.Acceptada = false
	p.Create()
	c.Data["json"] = p
	c.ServeJSON()

}

// PutTransaccio get a user
func (c *UserController) PutTransaccio() {
	//fer una peticio especifica a un usuari
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)

	var datapoint models.Peticio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	name := datapoint.Name
	description := datapoint.Descripcio

	userto := datapoint.To
	itemId := datapoint.ItemID
	uTo, _ := models.FindUserByID(userto)
	uPet, _ := models.FindUserByID(iduser)
	var pet models.Peticio
	pet.ID = EncodeMsg(iduser + time.Now().String())
	pet.Descripcio = description
	pet.IDuser = iduser
	pet.To = userto
	pet.Name = name
	pet.X = uPet.X
	pet.Y = uPet.Y
	pet.ItemID = itemId
	pet.Acceptada = true
	uTo.PutTransaccio(pet)
	uPet.PutTransaccio(pet)
	c.Data["json"] = "ok"
	c.ServeJSON()

}

type AcceptStruct struct {
	IDpet string `bson:"idpet"`
	IDit  string `bson:"idit"`
}

// AcceptRadiPetition put peticio al radi
func (c *UserController) AcceptRadiPetition() {
	//rebre el token i verificar si es coorrecte

	var datapoint AcceptStruct
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	idpet := datapoint.IDpet
	iditem := datapoint.IDit
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)

	fmt.Println(idpet)
	p, err := models.FindPeticioByID(idpet)

	if err != nil {
		c.Data["json"] = "Peticio ja acceptada"
	} else {
		p.To = iduser
		p.ItemID = iditem
		p.Acceptada = true
		models.DeletePeticioByID(idpet)
		uTo, _ := models.FindUserByID(p.To)
		uPet, _ := models.FindUserByID(p.IDuser)
		uTo.PutTransaccio(p)
		uPet.PutTransaccio(p)
		c.Data["json"] = "ok"

	}
	c.ServeJSON()
}

// GetPeticionsRadiUser get a user
func (c *UserController) GetPeticionsRadiUser() {
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		peticions, err := models.GetPeticionsRadi(u.X, u.Y, iduser)
		if err == nil {
			c.Data["json"] = peticions
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// GetPeticionsSelf get a user
func (c *UserController) GetPeticionsSelf() {
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	if err == nil {
		peticions, err := models.GetPeticionsSelf(iduser)
		if err == nil {
			c.Data["json"] = peticions
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// PutFavourite put a favourite to a user
func (c *UserController) PutFavourite() {
	iditem := c.GetString("idItem")
	idowner := c.GetString("idowner")
	token := c.Ctx.Input.Header("token")
	idusuari, err := DecodeToken(token)
	//buscar owner
	o, err := models.FindUserByID(idowner)
	//buscar objecte dins owner
	item, err := o.FindFavouriteByID(iditem)
	//put objecte a usuari
	u, err := models.FindUserByID(idusuari)
	if err != nil {
		c.Data["json"] = "error user not found"
	} else {
		u.PutFavouriteModel(item, idowner)
		c.Data["json"] = "ok"
	}
	c.ServeJSON()
}

//GetFavouritesUsuari get the user favourites
func (c *UserController) GetFavouritesUsuari() {
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		c.Data["json"] = u.FavUser
		c.ServeJSON()

	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}
