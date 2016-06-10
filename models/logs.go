package models

import (
	"sharit-backend/models/mongo"

	"gopkg.in/mgo.v2/bson"

	"github.com/astaxie/beego"
)

//"github.com/novikk/redline/models/mongo"

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lon"`
}

//Peticions is a list of User
type Logs []Point

// Create load the item to th db
func (p *Point) Create() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("logs")
	err := c.Insert(p)

	return err
}

func GetAllLogs() (Logs, error) {
	db := mongo.Conn()
	defer db.Close()
	var p Logs
	c := db.DB(beego.AppConfig.String("database")).C("logs")
	err := c.Find(bson.M{}).All(&p)
	return p, err
}
