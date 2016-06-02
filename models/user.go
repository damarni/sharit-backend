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
	ID                 bson.ObjectId `bson:"_id,omitempty"`
	IDuser             string        `bson:"iduser,omitempty"`
	Email              string        `bson:"email,omitempty"`
	Pass               string        `bson:"pass,omitempty"`
	Idioma             string        `bson:"idioma,omitempty"`
	Radi               float64       `bson:"radi,omitempty"`
	RadiReal           float64       `bson:"radireal,omitempty"`
	Name               string        `bson:"name,omitempty"`
	Surname            string        `bson:"surname,omitempty"`
	Stars              float64       `bson:"stars,omitempty"`
	ItemsUser          Items         `bson:"itemsUser,omitempty"`
	X                  float64       `bson:"x,omitempty"`
	Y                  float64       `bson:"y,omitempty"`
	Token              string        `bson:"token,omitempty"`
	FavUser            Favs          `bson:"favuser,omitempty"`
	Transaccions       []Peticio     `bson:"transaccions,omitempty"`
	Image              string        `bson:"image,omitempty"`
	NumeroAnuncios     int           `bson:"nmeroanuncios,omitempty"`
	NumeroPeticiones   int           `bson:"numeropeticiones,omitempty"`
	NumeroValoraciones int           `bson:"numerovaloracions,omitempty"`
	NumeroLikes        int           `bson:"nuemrolikes,omitempty"`
	NumeroPrestados    int           `bson:"numeroprestados"`
	NumeroPedidos      int           `bson:"numeropedidos:omitempty"`
	Valoracions        Vals          `bson:"valoracions"`
}

//numero de prestados y numero de pedidos se suma cuando se hace la valoracion
//aclarar que peticiones son, in o out, y si son pendientes, o si son todas

// UpdateNumeroLikes updates user profile
func (u *User) UpNumeroLikes() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$inc": bson.M{"quantity": 1, "NumeroLikes": 1}})
	return err
}

// DownNumeroLikes updates user profile
func (u *User) DownNumeroLikes() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$inc": bson.M{"quantity": -1, "NumeroLikes": 1}})
	return err
}

//Users is a list of User
type Users []User

// Create creates a user with its information in the database
func (u *User) Create() error {
	db := mongo.Conn()
	defer db.Close()
	fmt.Println("----------")

	fmt.Println(u.Radi)
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

// DeleteUserByID returns a user found by steamid
func DeleteUserByID(id string) error {

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Remove(bson.M{"iduser": id})

	return err
}

// FindUserByMail returns a user found by steamid
func FindUserByMail(mail string) (User, error) {
	var u User

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{"email": mail}).One(&u)

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

// UpdateItemModels updates user profile
func (u *User) UpdateItemModels(i Item) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"itemsUser": bson.M{"$elemMatch": bson.M{"Idd": i.Idd}, "iduser": u.IDuser}}, bson.M{"$set": bson.M{"itemsUser.$.itemname": i.ItemName, "itemsUser.$.description": i.Description, "itemsUser.$.image1": i.Image1}})
	return err
}

// PutComplainModel updates user profile
func (u *User) PutComplainModel(i string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"itemsUser": bson.M{"$elemMatch": bson.M{"Idd": i}, "iduser": u.IDuser}}, bson.M{"$inc": bson.M{"quantity": 1, "itemsUser.$.complains": 1}})
	return err
}

// UpdateUser updates user profile
func (u *User) UpdateUser() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"email": u.Email, "x": u.X, "y": u.Y, "radi": u.Radi, "idioma": u.Idioma, "radireal": u.RadiReal, "name": u.Name, "surname": u.Surname, "image": u.Image}})
	return err
}

// UpdateStars updates user profile
func (u *User) UpdateStars(stars float64) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"stars": stars}})
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

// PutTransaccio put item on a user array
func (u *User) PutTransaccio(p Peticio) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	fmt.Print("Models user")
	fmt.Println(u.IDuser)
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"transaccions": p}})

	fmt.Println(err)
	fmt.Println("-----")
	return err
}

// PutItemModel put item on a user array
func (u *User) PutItemModel(i Item) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	fmt.Println(u.IDuser)
	fmt.Println(i)
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"itemsUser": i}})
	fmt.Println(err)

	return err
}

// PutValoracio put item on a user array
func (u *User) PutValoracio(v Valoracio) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	fmt.Println("id user putval")
	fmt.Println(u.IDuser)
	fmt.Println("val")
	fmt.Println(v)
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"valoracions": v}})
	fmt.Println(err)
	return err
}

// DeleteItemModell put item on a user array
func (u *User) DeleteItemModel(id string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$pull": bson.M{"itemsUser": bson.M{"idd": id}}})
	fmt.Println(err)
	return err

}

// DeleteFavModell put item on a user array
func (u *User) DeleteFavModel(idItem, idUser string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$pull": bson.M{"favuser": bson.M{"iduser": idUser, "iditem": idItem}}})
	fmt.Println(err)
	return err

}

//  DeleteTransaccioModel put item on a user array
func (u *User) DeleteTransaccioModel(idTransacció string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	fmt.Println("id user delete trans")
	fmt.Println(u.IDuser)
	fmt.Println("id trans delete")
	fmt.Println(idTransacció)
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$pull": bson.M{"transaccions": bson.M{"id": idTransacció}}})
	fmt.Println(err)
	return err

}

// PutFavouriteModel put favourite on a user array FavUser
func (u *User) PutFavouriteModel(i, idowner string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	var f Fav
	f.IDuser = idowner
	f.IDitem = i
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"favuser": f}})
	return err
}

// GetUsersRadi returns a user found by steamid
func GetUsersRadi(x, y, radi float64) (Users, error) {
	var usrs Users

	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
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
		}}).All(&usrs)
	fmt.Println("----------------------")

	fmt.Println(usrs)
	fmt.Println("----------------------")

	return usrs, err
}

// GetItemsRadi returns a user found by steamid
func GetItemsRadi(x, y, radi float64) (Items, error) {
	var itms Items
	fmt.Println(x)
	fmt.Println(y)

	usrs, err := GetUsersRadi(x, y, radi)
	if err != nil {
		fmt.Println("error al get items")
	} else {
		for _, usr := range usrs {
			itms = append(itms, usr.ItemsUser...)
		}
		return itms, err
	}
	return itms, err

}
