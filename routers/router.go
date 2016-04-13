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

	beego.Router("/item/put", &controllers.ItemController{}, "get:Put")
	beego.Router("/item/getAll", &controllers.ItemController{}, "get:GetAll")
	beego.Router("/item/getAllRadi", &controllers.ItemController{}, "get:GetAllRadi")
}
