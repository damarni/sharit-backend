package routers

import (
	"sharitback/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/All", &controllers.MainController{}, "get:All")
}
