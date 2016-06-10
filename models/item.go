package models

import (
	"errors"
	"sharit-backend/models/mongo"
	"time"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	ID          uint64    `bson:"_id,omitempty"`
	Idd         string    `bson:"idd,omitempty"`
	IDuser      string    `bson:"iduser,omitempty"`
	ItemName    string    `bson:"itemname,omitempty"`
	Description string    `bson:"description,omitempty"`
	Image1      string    `bson:"image1,omitempty"`
	Image2      string    `bson:"image2,omitempty"`
	Image3      string    `bson:"image3,omitempty"`
	Stars       string    `bson:"stars,omitempty"`
	LastSharit  time.Time `bson:"lastSharit,omitempty"`
	Complains   int       `bson:"complains,omitempty"`
}

type Items []Item

func GetAllItems() (Items, error) {
	db := mongo.Conn()
	defer db.Close()
	var p Items
	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Find(bson.M{}).All(&p)
	return p, err
}

func (p *Item) FindByID(id string) error {
	db := mongo.Conn()
	defer db.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}

	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(p)

	return err
}

func (p *Item) Create() error {
	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Insert(p)

	return err
}
