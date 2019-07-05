package main

import (
	"admin/common"
	_ "admin/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)




func init()  {
	common.Init()


}


func AssetsDomains()(out string){
	return beego.AppConfig.String("AssetsDomains")
}


func main() {
	beego.SetStaticPath("/assets", "assets")

	//注册模板函数，资源域名独立配置
	beego.AddFuncMap("AssetsDomains",AssetsDomains)



	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

