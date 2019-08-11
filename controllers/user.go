package controllers

import (
	"admin/common"
	"admin/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"regexp"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
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
	this.Data["pageTitle"] = "添加账号"


	if this.isPost() {
		valid := validation.Validation{}

		username := this.GetString("username")
		email := this.GetString("email")
		sex, _ := this.GetInt("sex")
		status, _ := this.GetInt("status", 0)
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")

		valid.Required(username, "username").Message("请输入用户名")
		valid.Required(email, "email").Message("请输入Email")
		valid.Email(email, "email").Message("Email无效")
		valid.Required(password1, "password1").Message("请输入密码")
		valid.Required(password2, "password2").Message("请输入确认密码")
		valid.MinSize(password1, 6, "password1").Message("密码长度不能小于6个字符")
		valid.Match(password1, regexp.MustCompile(`^`+regexp.QuoteMeta(password2)+`$`), "password2").Message("两次输入的密码不一致")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				this.showMsg(err.Message, MSG_ERR)
			}
		}

		Id, _ := this.GetInt("Id",0)
		if Id == 0{
			//insert
			_, err := common.UserService.AddUser(username, email, password1, sex, status)
			this.checkError(err)
		}else{
			//update
			user := &common.User{Id:Id, UserName:username, Email:email, Password:common.GetPassowrdMd5Str(password1), Status:status, Sex:sex, UpdateTime:time.Now()}

			err := common.UserService.UpdateUser(user, "UserName", "Email", "Password", "Status", "Sex", "UpdateTime")
			this.checkError(err)
		}



		this.redirect(beego.URLFor("UserController.List"))
	}

	this.display()

}

// @router /user/edit [get,post]
func (this *UserController)Edit()  {
	this.Data["pageTitle"] = "编辑账号"
	this.TplName = "user/add.html"


	valid := validation.Validation{}

	userId, _ := this.GetInt("id",0)
	valid.Required(userId, "userId").Message("异常请求")
	valid.Min(userId,  1,"userId").Message("异常请求")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.showMsg(err.Message, MSG_ERR)
		}
	}


	user, err := common.UserService.GetUserById(userId)
	this.checkError(err)

	this.Data["user"] = user

	this.display()

}


// @router /user/Del [get]
func (this *UserController)Del()  {
	this.TplName = "user/list.html"



	valid := validation.Validation{}

	userId, _ := this.GetInt("id",0)
	valid.Required(userId, "userId").Message("异常请求")
	valid.Min(userId,  2,"userId").Message("管理员账号不能删除")
	if valid.HasErrors() {

		page := 1
		count, _ := common.UserService.GetTotal()
		users, _ := common.UserService.GetUserList(page, this.pageSize)
		this.Data["pageTitle"] = "帐号管理"
		this.Data["count"] = count
		this.Data["list"] = users
		this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("UserController.List"), true).ToString()

		for _, err := range valid.Errors {
			this.showMsg(err.Message, MSG_ERR)
		}
	}

	_, err := common.UserService.DeleteById(userId)
	this.checkError(err)

	this.redirect(beego.URLFor("UserController.List"))
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

	this.TplName = "user/login.html"

}



// @router /user/logout [get]
func (this *UserController) Logout() {

	this.DelSession("UserID")
	this.DelSession("UserName")
	this.redirect(beego.URLFor("UserController.Login"))
}
