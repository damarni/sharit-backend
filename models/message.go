package models

import (
	//"github.com/novikk/redline/models/mongo"

	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// Message is a user :D
type Message struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	First bool          `bson:"first,omitempty"`
	Text  string        `bson:"text,omitempty"`
}

//Messages is a list of User
type Messages []Message

// Create creates a user with its information in the database
func (p *Message) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("messages")
	err = c.Insert(p)
	return err
}
