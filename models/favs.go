package models

import "gopkg.in/mgo.v2/bson"

// Fav is a product :D
type Fav struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	IDuser string        `bson:"iduser,omitempty"`
	IDitem string        `bson:"iditem,omitempty"`
}

// Favs is a list of item
type Favs []Fav
