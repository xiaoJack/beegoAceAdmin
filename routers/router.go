package routers

import (
	"github.com/xiaoJack/beegoAceAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.ProjectController{})
	beego.Include(&controllers.ApiController{})
}
