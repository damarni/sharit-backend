package models

import (
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserId string        `bson:"userid,omitempty"`
	Text   string        `bson:"text,omitempty"`
	Date   string        `bson:"date,omitempty"`
}

type Messages []Message

func (p *Message) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("messages")
	err = c.Insert(p)
	return err
}
