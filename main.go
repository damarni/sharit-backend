package main

import (
	"sharit-backend/chat"
	_ "sharit-backend/routers"

	"github.com/astaxie/beego"
)

func main() {
	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// // - Credentials share
	go chat.Run()
	beego.Run()

}
