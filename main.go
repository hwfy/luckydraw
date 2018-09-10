package main

import (
	_ "luckyDraw/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/v1/luckydraw/apidoc"] = "swagger"
	}
	beego.Run()
}
