package models

import (
	//"github.com/novikk/redline/models/mongo"

	"errors"
	"fmt"
	"sharit-backend/models/mongo"
	"strconv"

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
	fmt.Println(u)
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
func (u *User) FindFavouriteByID(iditem string) (Item, error) {
	var itemaux Item
	var err error
	for _, fav := range u.ItemsUser {
		if strconv.Itoa(int(fav.ID)) == iditem {
			itemaux = fav
			return fav, nil
		}
	}
	err = errors.New("no item found")
	return itemaux, err

}

// UpdateUser updates user profile
func (u *User) UpdateUser() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"email": u.Email, "x": u.X, "y": u.Y}})
	return err
}

// UpdateUserCoords updates user cordenades
func (u *User) UpdateUserCoords() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"x": u.X, "y": u.Y}})
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

// PutPeticio updates user profile
func (u *User) PutPeticio(pet Peticio) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"peticions": pet}})
	return err
}

// PutFavouriteModel put favourite on a user array FavUser
func (u *User) PutFavouriteModel(i Item, idowner string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	var f Fav
	f.IDuser = idowner
	f.IDitem = strconv.Itoa(int(i.ID))
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"favuser": f}})
	return err
}

// GetUsersRadi returns a user found by steamid
func GetUsersRadi(x, y int) (Users, error) {
	var usrs Users

	db := mongo.Conn()
	defer db.Close()
	radi, _ := beego.AppConfig.Int("radi")
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(
		bson.M{"x": bson.M{"$lt": x + radi}},
	).All(&usrs)
	fmt.Println(usrs)
	return usrs, err
}

// GetItemsRadi returns a user found by steamid
func GetItemsRadi(x, y int) (Items, error) {
	var itms Items

	usrs, err := GetUsersRadi(x, y)
	if err == nil {
		fmt.Println("error al get items")
	} else {
		for _, usr := range usrs {
			itms = append(itms, usr.ItemsUser...)
		}
		return itms, err
	}
	return itms, err

}
