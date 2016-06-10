package models

import (
	"fmt"
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

type Peticio struct {
	IDmongo     bson.ObjectId `bson:"_id,omitempty"`
	ID          string        `bson:"id,omitempty"`
	IDuser      string        `bson:"iduser,omitempty"`
	UserName    string        `bson:"username,omitempty"`
	UserSurname string        `bson:"usersurname,omitempty"`
	Name        string        `bson:"name,omitempty"`
	To          string        `bson:"to,omitempty"`
	Descripcio  string        `bson:"descripcio,omitempty"`
	ItemID      string        `bson:"itemID,omitempty"`
	X           float64       `bson:"x,omitempty"`
	Y           float64       `bson:"y,omitempty"`
	Acceptada   bool          `bson:"acceptada"`
	Image       string        `bson:"image,omitempty"`
}

type Peticions []Peticio

func (p *Peticio) Create() error {
	db := mongo.Conn()
	defer db.Close()
	var err error
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err = c.Insert(p)
	return err
}

func GetPeticionsRadi(x, y, radi float64, iduser string) (Peticions, error) {
	var pets Peticions

	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err := c.Find(
		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{
					"$and": []interface{}{
						bson.M{"x": bson.M{"$lt": x + radi}},
						bson.M{"x": bson.M{"$gt": x - radi}}}},
				bson.M{
					"$and": []interface{}{
						bson.M{"y": bson.M{"$lt": x + radi}},
						bson.M{"y": bson.M{"$gt": x - radi}}}},
			}},
			bson.M{
				"$and": []interface{}{
					bson.M{"iduser": bson.M{"$ne": iduser}},
					bson.M{"acceptada": false}}},
		}}).All(&pets)
	return pets, err
}

func GetPeticionsSelf(iduser string) (Peticions, error) {
	var pets Peticions
	fmt.Println(iduser)
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err := c.Find(
		bson.M{
			"$and": []interface{}{
				bson.M{"to": iduser},
				bson.M{"acceptada": false}}}).All(&pets)
	return pets, err
}

func FindPeticioByID(id string) (Peticio, error) {
	var p Peticio

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err := c.Find(bson.M{"id": id}).One(&p)

	return p, err
}

func DeletePeticioByID(id string) error {

	db := mongo.Conn()
	defer db.Close()
	fmt.Println(id)
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err := c.Remove(bson.M{"id": id})

	return err
}

func (p *Peticio) UpdatePeticioTo() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("peticions")
	err := c.Update(bson.M{"id": p.ID}, bson.M{"$set": bson.M{"to": p.To, "itemID": p.ItemID}})
	return err
}
