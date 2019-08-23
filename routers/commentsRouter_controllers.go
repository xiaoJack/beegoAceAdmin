package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ApiController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ApiController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/api/add/:ProjectId:int/:LabelId:int`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ApiController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ApiController"],
        beego.ControllerComments{
            Method: "Details",
            Router: `/api/details/:ProjectId:int/:LabelId:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ApiController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ApiController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/api/edit/:ProjectId:int/:LabelId:int/:ApiId:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/project/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Addlabel",
            Router: `/project/addlabel/:Id:int`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Dellabel",
            Router: `/project/dellabel/:ProjectId:int/:LabelId:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Details",
            Router: `/project/details/:Id:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/project/edit/:Id:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Editlabel",
            Router: `/project/editlabel/:ProjectId:int/:LabelId:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/project/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Ztree",
            Router: `/project/ztree/:Id:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/user/Del`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/user/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/user/edit`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/user/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/user/login`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoJack/beegoAceAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/user/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
