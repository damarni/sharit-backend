package models

import (
	//"github.com/novikk/redline/models/mongo"

	"fmt"
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// Room is a user :D
type Room struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	RoomId  string        `bson:"roomid,omitempty"`
	UserID1 string        `bson:"userid1,omitempty"`
	UserID2 string        `bson:"userid2,omitempty"`
	IdTrans string        `bson:"idtrans,omitempty"`
	Rated1  bool          `bson:"rated1"`
	Rated2  bool          `bson:"rated2"`
	ItemID  string        `bson:"itemid,omitempty"`

	MessagesRoom Messages `bson:"messages,omitempty"`
}

//Rooms is a list of User
type Rooms []Room

// Create creates a user with its information in the database
func (r *Room) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("rooms")
	err = c.Insert(r)
	return err
}

//name users, name item, no missatges
// FindRooms returns a user found by steamid
func FindRooms(usid string) (Rooms, error) {
	var u Rooms
	var w Rooms
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("rooms")
	err := c.Find(bson.M{"userid1": usid}).All(&u)
	err = c.Find(bson.M{"userid2": usid}).All(&w)
	u = append(u, w...)
	return u, err
}

// FindRoom returns a user found by steamid
func FindRoom(id string) (Room, error) {
	var r Room
	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("rooms")
	err := c.Find(bson.M{"roomid": id}).One(&r)
	return r, err
}

// PutMessage put item on a user array
func (r *Room) PutMessage(i Message) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("rooms")
	err := c.Update(bson.M{"_id": r.ID}, bson.M{"$push": bson.M{"messages": i}})
	return err
}

// Rate put item on a user array
func (r *Room) Rate1() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("rooms")
	if r.Rated2 == true {
		fmt.Println("try remove")
		err := c.Remove(bson.M{"roomid": r.RoomId})
		return err
	} else {
		fmt.Println("try rate")
		err := c.Update(bson.M{"roomid": r.RoomId}, bson.M{"$set": bson.M{"rated1": true}})
		return err
	}
}

// Rate2 put item on a user array
func (r *Room) Rate2() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("rooms")
	if r.Rated1 == true {
		fmt.Println("try remove")
		err := c.Remove(bson.M{"roomid": r.RoomId})
		return err
	} else {
		fmt.Println("try rate")
		err := c.Update(bson.M{"roomid": r.RoomId}, bson.M{"$set": bson.M{"rated2": true}})
		return err
	}

}
