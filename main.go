package main

import (
	"github.com/hlinfocc/hlinfoPoh/etc"
	_ "github.com/hlinfocc/hlinfoPoh/routers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	etc.InitSys()
	beego.Run()
}
