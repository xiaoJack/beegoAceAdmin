package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}


// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("get,post", c.Login)
}


// @router /user/login [get,post]
func (this *UserController) Login() {

	if this.GetSession("UserID") != nil {
		this.redirect("/")
	}

	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		username := this.GetString("username")
		password := this.GetString("password")
		remember := this.GetString("remember")

		fmt.Print(remember)








		user, err := this.auth.Login(username, password)
		fmt.Print(err)

		if err != nil {
			this.Data["error"] = err.Error()
			this.TplName = "user/index.html"
		}else{
			this.SetSession("UserID", user.Id)
			this.SetSession("UserName", user.UserName)
			this.Redirect("/", 302)


		}

		//this.Data["json"] = data
		//this.ServeJSON()
		//this.StopRun()
	}else{
		this.TplName = "user/index.html"
	}

	
}



// @router /user/logout [get]
func (this *UserController) Logout() {

	this.DelSession("UserID")
	this.DelSession("UserName")
	this.redirect(beego.URLFor("UserController.Login"))
}
