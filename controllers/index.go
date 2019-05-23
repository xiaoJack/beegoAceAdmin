package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}


// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("get", c.Index)
	c.Mapping("get,post", c.Login)
}


// @router / [get]
func (this *IndexController) Index() {
	this.Data["pageTitle"] = "系统概况"
	this.display()
}


// @router /login [get,post]
func (this *IndexController) Login() {
	//if this.userId > 0 {
	//	this.redirect("/")
	//}

	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		//flash := beego.NewFlash()
		data := make(map[string]interface{}, 0)
		username := this.GetString("username")
		password := this.GetString("password")
		remember := this.GetString("remember")

		data["username"] = username
		data["password"] = password
		data["remember"] = remember
		this.Data["json"] = data
		this.ServeJSON()
		this.StopRun()
	}else{
		this.TplName = "index/login.html"
	}

	
}
