package controllers

import (
	"encoding/json"
	"fmt"
	"sharit-backend/models"
	"time"
)

type SocketController struct {
	BaseController
}

func (c *SocketController) CreateRoom() {
	var datapoint models.Room
	json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	fmt.Println("Rooms")
	fmt.Println(datapoint)
	fmt.Println("------")
	usid1 := datapoint.UserID1
	usid2 := datapoint.UserID2
	itemid := datapoint.ItemID
	fmt.Println(datapoint.IdTrans)
	var r models.Room
	r.UserID1 = usid1
	r.UserID2 = usid2
	r.ItemID = itemid
	r.Rated1 = false
	r.Rated2 = false
	r.IdTrans = datapoint.IdTrans
	aux := itemid + time.Now().String()
	auxID := EncodeID64(usid1, usid2, aux)
	r.RoomId = auxID
	err := r.Create()
	if err == nil {
		r, _ = models.FindRoom(auxID)
	}
	c.Data["json"] = r
	c.ServeJSON()
}

type roomWithUsers struct {
	RoomId    string
	UserID1   string
	UserID2   string
	ItemID    string
	NameU1    string
	NameU2    string
	SurnameU1 string
	SurnameU2 string
	NameItem  string
	IdTrans   string
	Rated1    bool
	Rated2    bool
}

type roomsWithUsers []roomWithUsers

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
		room.IdTrans = r.IdTrans
		room.Rated1 = r.Rated1
		room.Rated2 = r.Rated2
		room.NameU1 = u1.Name
		room.NameU2 = u2.Name
		room.SurnameU1 = u1.Surname
		room.SurnameU2 = u2.Surname
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

type GetRoomStruct struct {
	Room      models.Room
	NameU1    string
	NameU2    string
	SurnameU1 string
	SurnameU2 string
	NameItem  string
}

func (c *SocketController) GetRoom() {

	id := c.GetString("roomid")
	var room GetRoomStruct
	r, err := models.FindRoom(id)
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
	room.SurnameU1 = u1.Surname
	room.SurnameU2 = u2.Surname
	room.Room = r
	room.NameItem = item.ItemName
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = room
	}
	c.ServeJSON()

}
