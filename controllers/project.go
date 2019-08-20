package controllers

import (
	"github.com/xiaoJack/beegoAceAdmin/common"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ProjectController struct {
	BaseController
}



// @router /project/list [get]
func (this *ProjectController)List()  {

	this.Data["pageTitle"] = "API项目管理"
	var project common.Project


	list, _ := project.GetprojectList()

	this.Data["list"] = list
	this.display()
}



// @router /project/add [get,post]
func (this *ProjectController)Add()  {

	this.Data["pageTitle"] = "添加项目"
	if this.isPost() {

		var p  common.Project
		this.ParseForm(&p)
		//fmt.Println(project)
		b, err := p.Valid()
		if !b{
			this.Data["projectInfo"] = &p
			this.showMsg(err , MSG_ERR)
		}


		Id, _ := this.GetInt("Id",0)
		if Id == 0{
			//insert
			_, err := p.Add()
			this.checkError(err)
		}else{
			//update
			p.Id = Id
			err := p.UpdateProject("Project_name", "Project_describe", "Project_url", "Test_ip", "Release_ip", "Pro_ip", "Is_monitor", "Monitor_url")
			this.checkError(err)
		}

		this.redirect(beego.URLFor("ProjectController.List"))
	}

	this.display()
}



// @router /project/edit/:Id:int [get]
func (this *ProjectController)Edit()  {
	this.Data["pageTitle"] = "编辑项目"
	this.TplName = "project/add.html"
	var project common.Project

	valid := validation.Validation{}

	projectId := this.GetIntParamFromURL(":Id")

	valid.Required(projectId, "Id").Message("异常请求1")
	valid.Min(projectId,  1,"Id").Message("异常请求2")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.showMsg(err.Message, MSG_ERR)
		}
	}

	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)


	fmt.Println(projectInfo)
	this.Data["projectInfo"] = projectInfo

	this.display()

}