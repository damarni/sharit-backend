package controllers

import (
	"encoding/json"
	"sharit-backend/models"
	"time"
)

// SocketController does everything related to  login
type SocketController struct {
	BaseController
}

// CreateRoom register
func (c *SocketController) CreateRoom() {
	var datapoint models.Room
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	usid1 := datapoint.UserID1
	usid2 := datapoint.UserID2
	itemid := datapoint.ItemID
	var r models.Room
	r.UserID1 = usid1
	r.UserID2 = usid2
	r.ItemID = itemid
	aux := itemid + time.Now().String()
	r.RoomId = EncodeID64(usid1, usid2, aux)
	r.Create()
	c.Data["json"] = r
	c.ServeJSON()
}

type roomWithUsers struct {
	RoomId   string
	UserID1  string
	UserID2  string
	ItemID   string
	NameU1   string
	NameU2   string
	NameItem string
}

type roomsWithUsers []roomWithUsers

// GetRooms get a user
func (c *SocketController) GetRooms() {
	var retorn roomsWithUsers
	id := c.GetString("userid")

	u, err := models.FindRooms(id)
	for _, r := range u {
		var room roomWithUsers
		room.ItemID = r.ItemID
		room.RoomId = r.RoomId
		room.UserID1 = r.UserID1
		room.UserID2 = r.UserID2
		u1, _ := models.FindUserByID(r.UserID1)
		u2, _ := models.FindUserByID(r.UserID2)
		var item models.Item
		for _, it := range u2.ItemsUser {
			if it.Idd == r.ItemID {
				item = it
			}
		}
		room.NameU1 = u1.Name
		room.NameU2 = u2.Name
		room.NameItem = item.ItemName
		retorn = append(retorn, room)
	}
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = retorn
	}
	c.ServeJSON()

}

// GetRoom get a user
func (c *SocketController) GetRoom() {

	id := c.GetString("roomid")

	u, err := models.FindRoom(id)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// PutMessage get a user
/*func (c *SocketController) PutMessage() {

	idroom := c.GetString("idroom")
	text := c.GetString("text")
	user := c.GetString("user")
	var who bool
	who = false
	if user == "1" {
		who = true
	} else {
		who = false
	}
	var m models.Message
	m.First = who
	m.Text = text
	r, err := models.FindRoom(idroom)
	r.PutMessage(m)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = r
	}
	c.ServeJSON()

}*/
