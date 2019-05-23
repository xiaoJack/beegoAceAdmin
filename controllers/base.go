package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string               // 控制器名
	actionName     string               // 动作名
}


func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()

	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

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