package controllers

import "sharit-backend/models"

// SocketController does everything related to  login
type SocketController struct {
	BaseController
}

// CreateRoom register
func (c *SocketController) CreateRoom() {

	usid1 := c.GetString("userid1")
	udid2 := c.GetString("userid2")
	itemid := c.GetString("itemid")
	var r models.Room
	r.UserID1 = usid1
	r.UserID2 = usid2
	r.ItemID = itemid
	r.Create()
	c.Data["json"] = r
	c.ServeJSON()
}

// GetRooms get a user
func (c *SocketController) GetRooms() {

	id := c.GetString("userid")

	u, err := models.FindRooms(id)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// PutMessage get a user
func (c *SocketController) PutMessage() {

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
	r := models.FindRoom(idroom)
	r.PutMessage(m)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = r
	}
	c.ServeJSON()

}
