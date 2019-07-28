package controllers

import (
	"admin/common"
	"admin/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type UserController struct {
	BaseController
}


// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("get,post", c.Login)
}


// @router /user/list [get]
func (this *UserController)List()  {
	page, _ := strconv.Atoi(this.GetString("page"))
	if page < 1 {
		page = 1
	}

	count, _ := common.UserService.GetTotal()
	users, _ := common.UserService.GetUserList(page, this.pageSize)



	this.Data["pageTitle"] = "帐号管理"
	this.Data["count"] = count
	this.Data["list"] = users
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("UserController.List"), true).ToString()
	this.display()
	//this.TplName = "user/list.html"
}

// @router /user/add [get,post]
func (this *UserController)Add()  {
	this.TplName = "user/index.html"

}

// @router /user/edit [get,post]
func (this *UserController)Edit()  {
	this.TplName = "user/index.html"

}






// @router /user/login [get,post]
func (this *UserController) Login() {

	if this.GetSession("UserID") != nil {
		this.redirect("/")
	}


	//beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		isError := false
		username := this.GetString("username")
		password := this.GetString("password")
		remember := this.GetString("remember")

		fmt.Print(remember)

		valid := validation.Validation{}
		if v := valid.Required(username, "username"); !v.Ok {
			this.Data["error"] = "用户名不能为空"
			isError = true
		}

		if v := valid.Required(password, "password"); !v.Ok {
			this.Data["error"] = "密码不能为空"
			isError = true
		}



		if !isError {
			user, err := this.auth.Login(username, password)

			if err != nil {
				this.Data["error"] = err.Error()
			} else {
				this.SetSession("UserID", user.Id)
				this.SetSession("UserName", user.UserName)
				this.Redirect("/", 302)

			}
		}

	}

	this.TplName = "user/index.html"

}



// @router /user/logout [get]
func (this *UserController) Logout() {

	this.DelSession("UserID")
	this.DelSession("UserName")
	this.redirect(beego.URLFor("UserController.Login"))
}
