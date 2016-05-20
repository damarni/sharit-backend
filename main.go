package main

import (
	"sharit-backend/chat"
	_ "sharit-backend/routers"

	"github.com/astaxie/beego"
)

func main() {

	go chat.Run()
	beego.Run()

}
