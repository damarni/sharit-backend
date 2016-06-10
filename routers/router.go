package routers

import (
	"sharit-backend/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("*", &controllers.UserController{}, "options:Options")
	beego.Router("/user", &controllers.UserController{}, "post:Register")
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/users", &controllers.UserController{}, "get:GetAll")
	beego.Router("/user", &controllers.UserController{}, "get:Get")
	beego.Router("/user", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/user", &controllers.UserController{}, "put:EditProfile")
	beego.Router("/peticiones", &controllers.UserController{}, "get:GetPeticionsRadiUser")
	beego.Router("/peticionesSelf", &controllers.UserController{}, "get:GetPeticionsSelf")
	beego.Router("/peticion", &controllers.UserController{}, "post:PutPeticio")
	beego.Router("/peticion", &controllers.UserController{}, "delete:DeletePeticio")
	beego.Router("/acceptRadiPetition", &controllers.UserController{}, "put:AcceptRadiPetition")
	beego.Router("/anuncio", &controllers.UserController{}, "post:PutItem")
	beego.Router("/anuncio", &controllers.UserController{}, "get:GetItemSoft")
	beego.Router("/anuncio", &controllers.UserController{}, "put:UpdateItem")
	beego.Router("/anuncio", &controllers.UserController{}, "delete:DeleteItem")
	beego.Router("/itemsAll", &controllers.UserController{}, "get:GetItems")
	beego.Router("/anuncios", &controllers.UserController{}, "get:GetItemsRadi")
	beego.Router("/transaccions", &controllers.UserController{}, "get:GetTransaccions")
	beego.Router("/transaccion", &controllers.UserController{}, "post:PutTransaccio")
	beego.Router("/user", &controllers.UserController{}, "options:SendOptions")
	beego.Router("/room/create", &controllers.SocketController{}, "post:CreateRoom")
	beego.Router("/room/findRooms", &controllers.SocketController{}, "get:GetRooms")
	beego.Router("/room/findRoom", &controllers.SocketController{}, "get:GetRoom")
	beego.Router("/fav", &controllers.UserController{}, "post:PutFavourite")
	beego.Router("/favs", &controllers.UserController{}, "get:GetFavouritesUsuari")
	beego.Router("/fav", &controllers.UserController{}, "delete:DeleteFav")
	beego.Router("/valorarItem", &controllers.UserController{}, "post:ValorarItem")
	beego.Router("/valorarUser", &controllers.UserController{}, "post:ValorarUser")
	beego.Router("/valoracions", &controllers.UserController{}, "get:GetValoracions")
	beego.Router("/complain", &controllers.UserController{}, "post:PutComplain")
	//beego.Router("/room/putMessage", &controllers.SocketController{}, "get:PutMessage")

	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/dashboardFake", &controllers.DashboardFakeController{})
	//beego.Router("/user/putFavourite", &controllers.ItemController{}, "get:PutFavourite")
	//falta getFavourite
	//beego.Router("/user/putCoordenades", &controllers.ItemController{}, "get:PutCoordenades")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Accept", "Token", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
