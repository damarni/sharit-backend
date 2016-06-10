package controllers

import (
	"encoding/json"
	"sharit-backend/models"
	"time"
)

type UserController struct {
	BaseController
}

type LoginStruct struct {
	Email   string  `bson:"email"`
	X       float64 `bson:"x"`
	Y       float64 `bson:"y"`
	Pass    string  `bson:"pass"`
	Radi    float64 `bson:"radi"`
	Idioma  string  `bson:"radi"`
	Image   string  `bson:"image"`
	Name    string  `bson:"name"`
	Surname string  `bson:"surname"`
}

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

func (c *UserController) SendOptions() {

	token := c.Ctx.Input.Header("token")
	idToken, err := DecodeToken(token)

	if err == nil {

		err = models.DeleteUserByID(idToken)

		if err != nil {
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

func (c *UserController) Register() {
	var datapoint models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	name := datapoint.Name
	surname := datapoint.Surname
	stars := 0.0
	mail := datapoint.Email
	pass := datapoint.Pass
	image := datapoint.Image
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
		u.RadiReal = 0
		radi := 50.0
		u.Radi = radi
		radi = ((radi / 1000) / 6378) * (180 * 3.141592)
		u.RadiReal = radi
		u.Idioma = "es-ES"
		u.X = datapoint.X
		u.Y = datapoint.Y
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

func (c *UserController) Options() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.ServeJSON()
}

func (c *UserController) EditProfile() {
	var datapoint LoginStruct
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	mail := ""
	mail = datapoint.Email
	idioma := ""
	name := datapoint.Name
	surname := datapoint.Surname
	idioma = datapoint.Idioma
	radi := datapoint.Radi

	token := c.Ctx.Input.Header("token")
	id, err := DecodeToken(token)
	if err != nil {
		c.Data["json"] = "error token id"
		c.ServeJSON()
	}
	image := ""
	image = datapoint.Image

	coordx := -1.0
	coordy := -1.0
	coordx = datapoint.X
	coordy = datapoint.Y
	u, _ := models.FindUserByID(id)
	if mail != "" {
		u.Email = mail
	}
	if radi != 0 {
		u.Radi = radi
		radi = ((radi / 1000) / 6378) * (180 * 3.141592)
		u.RadiReal = radi
	}
	if idioma != "" {
		u.Idioma = idioma
	}

	if name != "" {
		u.Name = name
	}

	if surname != "" {
		u.Surname = surname
	}
	if coordx != -1 {
		u.X = coordx
	}
	if coordy != -1 {
		u.Y = coordy
	}
	if image != "" {
		u.Image = image
	}
	err = u.UpdateUser()
	if err != nil {
	} else {
	}
	c.Data["json"] = u
	c.ServeJSON()
}

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

func (c *UserController) DeleteUser() {

	token := c.Ctx.Input.Header("token")
	idToken, err := DecodeToken(token)

	if err == nil {

		err = models.DeleteUserByID(idToken)

		if err != nil {
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

func (c *UserController) DeletePeticio() {
	var datapoint models.Peticio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	token := c.Ctx.Input.Header("token")
	_, err := DecodeToken(token)
	idpet := datapoint.ID
	if err == nil {

		err = models.DeletePeticioByID(idpet)
		if err != nil {
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
	i.Complains = 0
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
		u.PutItemModel(i)
		u, _ := models.FindUserByID(iduser)
		c.Data["json"] = u
	}
	c.ServeJSON()

}

type Complain struct {
	IDuser string `bson:"iduser"`
	IDitem string `bson:"iditem"`
}

func (c *UserController) PutComplain() {
	var datapoint Complain
	token := c.Ctx.Input.Header("token")
	_, err := DecodeToken(token)
	u, err := models.FindUserByID(datapoint.IDuser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		u.PutComplainModel(datapoint.IDitem)
		u, err = models.FindUserByID(datapoint.IDuser)
		i := c.GetItem(datapoint.IDitem, datapoint.IDuser)
		if i.Complains == 10 {
			u.DeleteItemModel(datapoint.IDitem)
			c.Data["json"] = "Item deleted"
		} else {
			c.Data["json"] = "Complain ok"
		}
	}
	c.ServeJSON()

}

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
		u.DeleteItemModel(idItem)
		u, _ := models.FindUserByID(iduser)
		c.Data["json"] = u
	}
	c.ServeJSON()

}

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

func (c *UserController) GetValoracions() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u.Valoracions
	}
	c.ServeJSON()

}

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

func (c *UserController) GetItemsRadi() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		items, err := models.GetItemsRadi(u.X, u.Y, u.RadiReal)
		if err == nil {
			c.Data["json"] = items
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

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

func (c *UserController) GetItem(idItem, idUser string) models.Item {

	u, err := models.FindUserByID(idUser)
	var item models.Item
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		items := u.ItemsUser
		for _, it := range items {
			if it.Idd == idItem {
				item = it
			}
		}
	}
	return item
}

type GetItemSoftStruct struct {
	IDuser  string  `bson:"iduser,omitempty"`
	Name    string  `bson:"name,omitempty"`
	Surname string  `bson:"surname,omitempty"`
	Stars   float64 `bson:"stars,omitempty"`
	X       float64 `bson:"x,omitempty"`
	Y       float64 `bson:"y,omitempty"`
	It      models.Item
}

func (c *UserController) GetItemSoft() {
	token := c.Ctx.Input.Header("token")
	idProp, err := DecodeToken(token)
	idUser := c.GetString("idUser")
	if idUser == "" {
		idUser = idProp
	}
	idItem := c.GetString("idItem")
	u, err := models.FindUserByID(idUser)
	var ret GetItemSoftStruct
	ret.Name = u.Name
	ret.Surname = u.Surname
	ret.IDuser = u.IDuser
	ret.Stars = u.Stars
	ret.X = u.X
	ret.Y = u.Y
	var item models.Item
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		items := u.ItemsUser
		for _, it := range items {
			if it.Idd == idItem {
				item = it
			}
		}
		ret.It = item
	}
	c.Data["json"] = ret

	c.ServeJSON()
}

func (c *UserController) PutPeticio() {
	token := c.Ctx.Input.Header("token")
	iduser, _ := DecodeToken(token)

	var datapoint models.Peticio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	name := datapoint.Name
	description := datapoint.Descripcio
	image := datapoint.Image
	u, _ := models.FindUserByID(iduser)
	var p models.Peticio
	p.IDuser = ""
	p.UserName = u.Name
	p.UserSurname = u.Surname
	p.Image = image
	p.ID = EncodeMsg(iduser + time.Now().String())
	p.Name = name
	p.To = iduser
	p.Descripcio = description
	p.X = u.X
	p.Y = u.Y
	p.Acceptada = false
	p.Create()
	c.Data["json"] = p
	c.ServeJSON()

}

type ReturnTrans struct {
	IDTrans string `bson:"idtrans,omitempty"`
}

func (c *UserController) PutTransaccio() {
	token := c.Ctx.Input.Header("token")
	userto, _ := DecodeToken(token)

	var datapoint models.Peticio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)

	name := datapoint.Name
	description := datapoint.Descripcio

	iduser := datapoint.IDuser
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
	var t ReturnTrans
	t.IDTrans = pet.ID
	c.Data["json"] = t
	var p models.Point
	p.Lat = uPet.X
	p.Lng = uPet.Y
	p.Create()
	c.ServeJSON()

}

type AcceptStruct struct {
	IDpet string `bson:"idpet"`
	IDit  string `bson:"idit"`
}

func (c *UserController) AcceptRadiPetition() {
	var datapoint AcceptStruct
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	idpet := datapoint.IDpet
	iditem := datapoint.IDit
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
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
		var p models.Point
		p.Lat = uPet.X
		p.Lng = uPet.Y
		c.Data["json"] = "ok"
	}
	c.ServeJSON()
}

type ValoracioCall struct {
	IDpet     string  `bson:"idpet"`
	Valoracio string  `bson:"valoracio"`
	Stars     float64 `bson:"stars"`
	User      string  `bson:"user"`
	IDitem    string  `bson:"iditem,omitempty"`
	RoomId    string  `bson:"roomid,omitempty"`
	Name      string  `bson:"name,omitempty"`
	Surname   string  `bson:"surname,omitempty"`
}

func (c *UserController) ValorarItem() {
	var datapoint ValoracioCall
	var val models.Valoracio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	val.IDitem = datapoint.IDitem
	val.IDtrans = datapoint.IDpet
	val.Stars = datapoint.Stars
	val.Valoracio = datapoint.Valoracio
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	val.Name = u.Name
	val.Surname = u.Surname
	val.User = u.IDuser
	user, err := models.FindUserByID(datapoint.User)
	if err != nil {
		c.Data["json"] = "Peticio ja acceptada"
	} else {
		x := user.Stars
		y := float64(len(user.Valoracions))
		new := ((x * y) + datapoint.Stars) / (y + 1)
		user.UpdateStars(new)
		user.PutValoracio(val)
		u.DeleteTransaccioModel(datapoint.IDpet)
		room, _ := models.FindRoom(datapoint.RoomId)
		room.Rate1()
		c.Data["json"] = "ok"
	}
	c.ServeJSON()
}

func (c *UserController) ValorarUser() {
	var datapoint ValoracioCall
	var val models.Valoracio
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	val.IDitem = datapoint.IDitem
	val.IDtrans = datapoint.IDpet
	val.Stars = datapoint.Stars
	val.Valoracio = datapoint.Valoracio
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	val.Name = u.Name
	val.Surname = u.Surname
	val.User = u.IDuser
	user, err := models.FindUserByID(datapoint.User)
	if err != nil {
		c.Data["json"] = "Peticio ja acceptada"
	} else {
		x := user.Stars
		y := float64(len(user.Valoracions))
		new := ((x * y) + datapoint.Stars) / (y + 1)
		user.UpdateStars(new)
		user.PutValoracio(val)
		u.DeleteTransaccioModel(datapoint.IDpet)
		room, _ := models.FindRoom(datapoint.RoomId)
		room.Rate2()
		c.Data["json"] = "ok"

	}
	c.ServeJSON()
}

func (c *UserController) GetPeticionsRadiUser() {
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		peticions, err := models.GetPeticionsRadi(u.X, u.Y, u.RadiReal, iduser)
		if err == nil {
			c.Data["json"] = peticions
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

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

func (c *UserController) PutFavourite() {
	var datapoint models.Fav
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	o, _ := models.FindUserByID(datapoint.IDuser)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "error user not found"
	} else {
		u.PutFavouriteModel(datapoint.IDitem, datapoint.IDuser)
		c.Data["json"] = "ok"
		o.UpNumeroLikes()
	}
	c.ServeJSON()
}

func (c *UserController) GetFavouritesUsuari() {
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		var its models.Items
		for _, fav := range u.FavUser {
			its = append(its, c.GetItem(fav.IDitem, fav.IDuser))
		}
		c.Data["json"] = its
		c.ServeJSON()

	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

func (c *UserController) DeleteFav() {
	var datapoint models.Fav
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	token := c.Ctx.Input.Header("token")
	iduser, err := DecodeToken(token)
	o, err := models.FindUserByID(datapoint.IDuser)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		u.DeleteFavModel(datapoint.IDitem, datapoint.IDuser)
		u, _ := models.FindUserByID(iduser)
		o.DownNumeroLikes()
		c.Data["json"] = u
	}
	c.ServeJSON()

}
