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

func (u *User) UpNumeroLikes() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$inc": bson.M{"quantity": 1, "NumeroLikes": 1}})
	return err
}

func (u *User) DownNumeroLikes() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$inc": bson.M{"quantity": -1, "NumeroLikes": 1}})
	return err
}

type Users []User

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

func FindUserByID(id string) (User, error) {
	var u User

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{"iduser": id}).One(&u)

	return u, err
}

func DeleteUserByID(id string) error {

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Remove(bson.M{"iduser": id})

	return err
}

func FindUserByMail(mail string) (User, error) {
	var u User

	db := mongo.Conn()
	defer db.Close()

	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{"email": mail}).One(&u)

	return u, err
}

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

func (u *User) UpdateItemModels(i Item) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"itemsUser": bson.M{"$elemMatch": bson.M{"Idd": i.Idd}, "iduser": u.IDuser}}, bson.M{"$set": bson.M{"itemsUser.$.itemname": i.ItemName, "itemsUser.$.description": i.Description, "itemsUser.$.image1": i.Image1}})
	return err
}

func (u *User) PutComplainModel(i string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"itemsUser": bson.M{"$elemMatch": bson.M{"Idd": i}, "iduser": u.IDuser}}, bson.M{"$inc": bson.M{"quantity": 1, "itemsUser.$.complains": 1}})
	return err
}

func (u *User) UpdateUser() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"email": u.Email, "x": u.X, "y": u.Y, "radi": u.Radi, "idioma": u.Idioma, "radireal": u.RadiReal, "name": u.Name, "surname": u.Surname, "image": u.Image}})
	return err
}

func (u *User) UpdateStars(stars float64) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"stars": stars}})
	return err
}

func (u *User) UpdateUserCoords() error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$set": bson.M{"x": u.X, "y": u.Y}})
	return err
}

func GetAllUsers() (Users, error) {
	db := mongo.Conn()
	defer db.Close()
	var p Users
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Find(bson.M{}).All(&p)
	return p, err
}

func (u *User) PutTransaccio(p Peticio) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"transaccions": p}})
	return err
}

func (u *User) PutItemModel(i Item) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"itemsUser": i}})
	return err
}

func (u *User) PutValoracio(v Valoracio) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$push": bson.M{"valoracions": v}})
	return err
}

func (u *User) DeleteItemModel(id string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$pull": bson.M{"itemsUser": bson.M{"idd": id}}})
	return err

}

func (u *User) DeleteFavModel(idItem, idUser string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$pull": bson.M{"favuser": bson.M{"iduser": idUser, "iditem": idItem}}})
	return err

}

func (u *User) DeleteTransaccioModel(idTransacció string) error {
	db := mongo.Conn()
	defer db.Close()
	c := db.DB(beego.AppConfig.String("database")).C("users")
	err := c.Update(bson.M{"iduser": u.IDuser}, bson.M{"$pull": bson.M{"transaccions": bson.M{"id": idTransacció}}})
	return err

}

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

	return usrs, err
}

func GetItemsRadi(x, y, radi float64) (Items, error) {
	var itms Items
	usrs, err := GetUsersRadi(x, y, radi)
	if err != nil {
	} else {
		for _, usr := range usrs {
			itms = append(itms, usr.ItemsUser...)
		}
		return itms, err
	}
	return itms, err

}
