package routers

import (
	"sharit-backend/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/All", &controllers.MainController{}, "get:All")
	beego.Router("/user/register", &controllers.UserController{}, "put:Register")
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/getAll", &controllers.UserController{}, "get:GetAll")
	beego.Router("/user/get", &controllers.UserController{}, "get:Get")
	beego.Router("/user/updateUser", &controllers.UserController{}, "post:EditProfile")
	beego.Router("/user/getPeticionsRadi", &controllers.UserController{}, "get:GetPeticionsRadiUser")
	beego.Router("/user/putPeticioRadi", &controllers.UserController{}, "put:PutPeticioRadi")
	beego.Router("/user/getPeticionsUsuari", &controllers.UserController{}, "get:GetPeticionsUsuari")
	beego.Router("/user/putPeticioUsuari", &controllers.UserController{}, "put:PutPeticioUsuari")
	beego.Router("/user/putItem", &controllers.UserController{}, "put:PutItem")
	beego.Router("/user/getItems", &controllers.UserController{}, "get:GetItems")
	beego.Router("/user/getItemsRadi", &controllers.UserController{}, "get:GetItemsRadi")

	beego.Router("/room/create", &controllers.SocketController{}, "get:CreateRoom")
	beego.Router("/room/findRooms", &controllers.SocketController{}, "get:GetRooms")
	beego.Router("/room/findRoom", &controllers.SocketController{}, "get:GetRoom")
	beego.Router("/room/putMessage", &controllers.SocketController{}, "get:PutMessage")

	//beego.Router("/user/putFavourite", &controllers.ItemController{}, "get:PutFavourite")
	//falta getFavourite
	//beego.Router("/user/putCoordenades", &controllers.ItemController{}, "get:PutCoordenades")

	var FilterCORS = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterCORS)
}
