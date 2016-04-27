package models

import (
	//"github.com/novikk/redline/models/mongo"

	"sharit-backend/models"
	"sharit-backend/models/mongo"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// User is a user :D
type User struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	IDuser        string        `bson:"iduser,omitempty"`
	Email         string        `bson:"email,omitempty"`
	Pass          string        `bson:"pass,omitempty"`
	Name          string        `bson:"name,omitempty"`
	Surname       string        `bson:"surname,omitempty"`
	Stars         string        `bson:"stars,omitempty"`
	ItemsUser     Items         `bson:"itemsUser,omitempty"`
	X             int           `bson:"x,omitempty"`
	Y             int           `bson:"y,omitempty"`
	Token         string        `bson:"token,omitempty"`
	FavUser       Favs          `bson:"favuser,omitempty"`
	PeticionsUser Peticions     `bson:"peticions,omitempty"`
}

//Users is a list of User
type Users []User

// Create creates a user with its information in the database
func (u *User) Create() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Insert(u)
	return err
}

// FindUserByID returns a user found by steamid
func FindUserByID(id string) (User, error) {
	var u User

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{"iduser": id}).One(&u)

	return u, err
}

//FindFavouriteByID returns de favourite with the id idItem
func FindFavouriteByID(iditem string, u *User) (Item, error) {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")

	err := c.Find(bson.M{"iduser": id}).One(&u)
	var i models.Item

	p := 0

	for( p < len(u.FavUser)) {
		if( u.FavUser[p].ID == iditem) i := u.FavUser[p]
		else ++p
	}

	return i, err

}

// UpdateUser updates user profile
func (u *User) UpdateUser() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"email": u.Email, "x": u.X, "y": u.Y}})
	return err
}

// GetAllUsers returns all users
func GetAllUsers() (Users, error) {
	db := mongo.Conn()
	defer db.Close()
	var p Users
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{}).All(&p)
	return p, err
}

// PutItemModel put item on a user array
func (u *User) PutItemModel(i Item) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"itemsUser": i}})
	return err
}

// PutFavouriteModel put favourite on a user array FavUser
func (u *User) PutFavouriteModel(i Item) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"favuser": i}})
	return err
}
