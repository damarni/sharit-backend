package routers

import (
	"sharit-backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/All", &controllers.MainController{}, "get:All")
	beego.Router("/user/register", &controllers.UserController{}, "get:Register")
	beego.Router("/user/getAll", &controllers.UserController{}, "get:GetAll")
	beego.Router("/user/get", &controllers.UserController{}, "get:Get")
	beego.Router("/user/updateUser", &controllers.UserController{}, "get:EditProfile")
	beego.Router("/user/registerdebug", &controllers.UserController{}, "get:RegisterDebug")
	beego.Router("/user/getPeticionsRadi", &controllers.UserController{}, "get:GetPeticionsRadiUser")
	beego.Router("/user/putPeticioRadi", &controllers.UserController{}, "get:PutPeticioRadi")
	beego.Router("/user/getPeticionsUsuari", &controllers.UserController{}, "get:GetPeticionsUsuari")
	beego.Router("/user/putPeticioUsuari", &controllers.UserController{}, "get:PutPeticioUsuari")
	beego.Router("/user/putItem", &controllers.UserController{}, "post:PutItem")
	beego.Router("/user/getItems", &controllers.UserController{}, "get:GetItems")

	beego.Router("/room/create", &controllers.SocketController{}, "get:CreateRoom")
	beego.Router("/room/findRooms", &controllers.SocketController{}, "get:GetRooms")
	beego.Router("/room/findRoom", &controllers.SocketController{}, "get:GetRoom")
	beego.Router("/room/putMessage", &controllers.SocketController{}, "get:PutMessage")

	beego.Router("/item/put", &controllers.ItemController{}, "get:Put")
	beego.Router("/item/getAll", &controllers.ItemController{}, "get:GetAll")
	beego.Router("/item/getAllRadi", &controllers.ItemController{}, "get:GetAllRadi")

	//beego.Router("/user/putFavourite", &controllers.ItemController{}, "get:PutFavourite")
	//falta getFavourite
	//beego.Router("/user/putCoordenades", &controllers.ItemController{}, "get:PutCoordenades")
}
