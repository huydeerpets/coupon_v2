package main

import (
	_ "ttc/cps/routers"

	"github.com/astaxie/beego"
	_ "ttc/cps/conf/inits"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

