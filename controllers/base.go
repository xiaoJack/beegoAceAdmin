package controllers

import (
	"admin/common"
	"github.com/astaxie/beego"
	"strings"
)


const (
	MSG_OK       = 0  // ajax输出错误码，成功
	MSG_ERR      = -1 // 错误
	MSG_REDIRECT = -2 // 重定向
)



type BaseController struct {
	beego.Controller
	controllerName string               // 控制器名
	actionName     string               // 动作名
	auth           *common.AuthService // 验证服务
}




func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()

	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	if this.actionName != "login" {
		this.initAuth()
	}


}



//登录验证
func (this *BaseController) initAuth() {
	userId := this.GetSession("UserID")

	if userId == nil {
		this.redirect(beego.URLFor("UserController.Login"))

	} else {
		//登入用户名字
		uname := this.GetSession("UserName")
		this.Data["userName"] = uname.(string)
	}
}



// 重定向
func (this *BaseController) redirect(url string) {
	if this.IsAjax() {
		this.showMsg("", MSG_REDIRECT, url)
	} else {
		this.Redirect(url, 302)
		this.StopRun()
	}
}


// 提示消息
func (this *BaseController) showMsg(msg string, msgno int, redirect ...string) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg
	out["redirect"] = ""
	if len(redirect) > 0 {
		out["redirect"] = redirect[0]
	}

	if this.IsAjax() {
		this.jsonResult(out)
	} else {
		for k, v := range out {
			this.Data[k] = v
		}
		this.display("error/message")
		this.Render()
		this.StopRun()
	}
}



// 输出json
func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}



// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}


//渲染模版
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}

	this.Layout = "layout/layout.html"
	this.TplName = tplname

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "layout/header.html"
	this.LayoutSections["Footer"] = "layout/footer.html"
	this.LayoutSections["Navbar"] = "layout/navbar.html"
	this.LayoutSections["Sidebar"] = "layout/sidebar.html"
}