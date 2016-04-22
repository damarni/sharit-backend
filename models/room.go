package models

import (
	//"github.com/novikk/redline/models/mongo"

	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// Room is a user :D
type Room struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	UserID1      string        `bson:"userid1,omitempty"`
	UserID2      string        `bson:"userid2,omitempty"`
	ItemID       string        `bson:"itemid,omitempty"`
	MessagesRoom Messages      `bson:"messages,omitempty"`
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
	err := c.Find(bson.M{"userid1": usid}).One(&r)
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
