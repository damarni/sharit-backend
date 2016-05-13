package routers

import (
	"sharit-backend/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "put:Register")
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/users", &controllers.UserController{}, "get:GetAll")
	beego.Router("/user", &controllers.UserController{}, "get:Get")
	beego.Router("/user", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/user", &controllers.UserController{}, "post:EditProfile")
	beego.Router("/peticiones", &controllers.UserController{}, "get:GetPeticionsRadiUser")
	beego.Router("/peticion", &controllers.UserController{}, "put:PutPeticio")
	beego.Router("/peticion", &controllers.UserController{}, "delete:DeletePeticio")
	beego.Router("/peticionesUsuario", &controllers.UserController{}, "get:GetPeticionsUsuari")
	beego.Router("/acceptRadiPetition", &controllers.UserController{}, "put:AcceptRadiPetition")
	beego.Router("/anuncio", &controllers.UserController{}, "put:PutItem")
	beego.Router("/anuncio", &controllers.UserController{}, "get:GetItemSoft")
	beego.Router("/anuncio", &controllers.UserController{}, "post:UpdateItem")
	beego.Router("/anuncio", &controllers.UserController{}, "delete:DeleteItem")
	beego.Router("/itemsAll", &controllers.UserController{}, "get:GetItems")
	beego.Router("/anuncios", &controllers.UserController{}, "get:GetItemsRadi")

	beego.Router("/room/create", &controllers.SocketController{}, "get:CreateRoom")
	beego.Router("/room/findRooms", &controllers.SocketController{}, "get:GetRooms")
	beego.Router("/room/findRoom", &controllers.SocketController{}, "get:GetRoom")
	beego.Router("/room/putMessage", &controllers.SocketController{}, "get:PutMessage")

	beego.Router("/dashboard", &controllers.DashboardController{})

	//beego.Router("/user/putFavourite", &controllers.ItemController{}, "get:PutFavourite")
	//falta getFavourite
	//beego.Router("/user/putCoordenades", &controllers.ItemController{}, "get:PutCoordenades")

	var FilterCORS = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterCORS)
}
