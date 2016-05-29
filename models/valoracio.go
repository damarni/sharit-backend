package models

import "gopkg.in/mgo.v2/bson"

// Fav is a product :D
type Valoracio struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	IDtrans   string        `bson:"idpet"`
	Valoracio string        `bson:"valoracio"`
	Stars     float64       `bson:"stars"`
	User      string        `bson:"user"`
	IDitem    string        `bson:"iditem,omitempty"`
}

// Favs is a list of item
type Vals []Valoracio
