package models

import (
	"errors"
	"sharit-backend/models/mongo"
	"time"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// Item is a product :D
type Item struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	ItemName    string        `bson:"item,omitempty"`
	Description string        `bson:"description,omitempty"`
	Image       string        `bson:"imagen,omitempty"`
	Stars       string        `bson:"stars,omitempty"`
	LastSharit  time.Time     `bson:"lastSharit,omitempty"`
	OwnerName   string        `bson:"ownerName,omitempty"`
	X           string        `bson:"x,omitempty"`
	Y           string        `bson:"y,omitempty"`
}

// Items is a list of item
type Items []Item

// GetAllItems returns all items
func GetAllItems() (Items, error) {
	db := mongo.Conn()
	defer db.Close()
	var p Items
	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Find(bson.M{}).All(&p)
	return p, err
}

// FindByID returns item given id
func (p *Item) FindByID(id string) error {
	db := mongo.Conn()
	defer db.Close()

	// check valid ID
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}

	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(p)

	return err
}

// Create load the item to th db
func (p *Item) Create() error {
	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Insert(p)

	return err
}
