package main

import (
	"github.com/hlinfocc/hlinfoPwh/etc"
	_ "github.com/hlinfocc/hlinfoPwh/routers"

	_ "github.com/lib/pq"
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
