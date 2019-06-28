package main

import (
	"admin/common"
	"github.com/astaxie/beego"
	_ "admin/routers"
)

func init()  {
	common.Init()


}



func main() {
	beego.SetStaticPath("/assets", "assets")
	beego.Run()
}

