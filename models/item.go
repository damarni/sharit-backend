package models

import (
	"errors"
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// Product is a product :D
type Product struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	ItemName string        `bson:"item,omitempty"`
	Type     string        `bson:"categoria,omitempty"`
	Image    string        `bson:"imagen,omitempty"`
}

// Products is a list of products
type Products []Product

// FindAll returns all products
func (p *Products) FindAll() error {
	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Find(bson.M{}).All(p)

	return err
}

// FindByID returns product given id
func (p *Product) FindByID(id string) error {
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
