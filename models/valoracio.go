package models

// Fav is a product :D
type Valoracio struct {
	IDtrans   string  `bson:"idpet"`
	Valoracio string  `bson:"valoracio"`
	Stars     float64 `bson:"stars"`
	User      string  `bson:"user"`
	IDitem    string  `bson:"iditem,omitempty"`
	Name      string  `bson:"name,omitempty"`
	Surname   string  `bson:"surname,omitempty"`
}

// Favs is a list of item
type Vals []Valoracio
