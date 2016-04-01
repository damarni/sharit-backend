package models

import (
	"errors"

	"sharitback/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// ProductType is a product type :D
type ProductType string

// Product Types
const (
	ProductKey   ProductType = "Keys"
	ProductCase  ProductType = "Cases"
	ProductSkin  ProductType = "Skins"
	ProductOther ProductType = "Others"
)

// Product is a product :D
type Product struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	ItemName string        `bson:"item,omitempty"`
	Type     ProductType   `bson:"categoria,omitempty"`
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

// FindAllProductsOfType returns all products filtered by type
func FindAllProductsOfType(ptype ProductType) (Products, error) {
	var p Products
	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("items")
	err := c.Find(bson.M{"categoria": ptype}).All(&p)

	return p, err
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
