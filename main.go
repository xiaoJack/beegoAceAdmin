package main

import (
	_ "admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/assets", "assets")
	beego.Run()
}

