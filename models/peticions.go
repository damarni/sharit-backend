package models

import (
	//"github.com/novikk/redline/models/mongo"

	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// Peticio is a user :D
type Peticio struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	IDuser     string        `bson:"iduser,omitempty"`
	Name       string        `bson:"name,omitempty"`
	To         string        `bson:"to,omitempty"`
	Descripcio string        `bson:"descripcio,omitempty"`
	X          int           `bson:"x,omitempty"`
	Y          int           `bson:"y,omitempty"`
}

//Peticions is a list of User
type Peticions []Peticio

// Create creates a user with its information in the database
func (p *Peticio) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err = c.Insert(p)
	return err
}

// GetPeticionsRadi returns a user found by steamid
func GetPeticionsRadi(x, y int) (Peticions, error) {
	var pets Peticions

	db := mongo.Conn()
	defer db.Close()
	radi, _ := beego.AppConfig.Int("radi")
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err := c.Find(
		bson.M{"$and": []interface{}{
			bson.M{
				"$and": []interface{}{
					bson.M{"x": bson.M{"$lt": x + radi}},
					bson.M{"x": bson.M{"$gt": x - radi}}}},
			bson.M{
				"$and": []interface{}{
					bson.M{"y": bson.M{"$lt": x + radi}},
					bson.M{"y": bson.M{"$gt": x - radi}}}},
		}}).All(&pets)
	return pets, err
}
