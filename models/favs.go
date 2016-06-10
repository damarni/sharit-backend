package models

import (
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type Fav struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	IDuser string        `bson:"iduser,omitempty"`
	IDitem string        `bson:"iditem,omitempty"`
}

type Favs []Fav

func (f *Fav) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("favorits")
	err = c.Insert(f)
	return err
}
