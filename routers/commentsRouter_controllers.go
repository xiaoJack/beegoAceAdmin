package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["admin/controllers:IndexController"] = append(beego.GlobalControllerRouter["admin/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["admin/controllers:UserController"] = append(beego.GlobalControllerRouter["admin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/user/login`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["admin/controllers:UserController"] = append(beego.GlobalControllerRouter["admin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/user/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
