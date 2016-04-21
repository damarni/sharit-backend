package models

import (
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// Fav is a product :D
type Fav struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	IDuser string        `bson:"iduser,omitempty"`
	IDitem string        `bson:"iditem,omitempty"`
}

// Favs is a list of item
type Favs []Fav

//Create creates a favourite with its information in the database
func (f *Fav) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("favorits")
	err = c.Insert(f)
	return err
}
