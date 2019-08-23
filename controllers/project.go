package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/xiaoJack/beegoAceAdmin/common"
	"github.com/xiaoJack/beegoAceAdmin/libs"
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



// @router /project/details/:Id:int [get]
func (this *ProjectController)Details()  {
	this.Data["pageTitle"] = "项目详情"
	var project common.Project


	projectId := this.GetIntParamFromURL(":Id")
	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)


	this.Data["projectInfo"] = projectInfo

	this.display()
}



// @router /project/addlabel/:Id:int [get,post]
func (this *ProjectController)Addlabel()  {
	this.Data["pageTitle"] = "增加标签"
	var project common.Project
	var projectLabel common.Project_label


	projectId := this.GetIntParamFromURL(":Id")
	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)

	if this.isPost() {
		this.ParseForm(&projectLabel)
		b, err := projectLabel.Valid()
		if !b{
			this.Data["projectInfo"] = projectInfo
			this.Data["projectLabel"] = &projectLabel
			this.showMsg(err , MSG_ERR)
		}


		Id, _ := this.GetInt("Id",0)
		if Id == 0{
			//insert
			_, err := projectLabel.Add()
			this.checkError(err)
		}else{
			//update
			projectLabel.Id = Id
			err := projectLabel.UpdateProject("Label_name",)
			this.checkError(err)
		}

		this.redirect(beego.URLFor("ProjectController.Details",":Id", projectId))

	}

	this.Data["projectInfo"] = projectInfo
	this.display()
}



// @router /project/editlabel/:ProjectId:int/:LabelId:int [get]
func (this *ProjectController)Editlabel()  {
	this.Data["pageTitle"] = "修改标签"
	var project common.Project
	var projectLabel common.Project_label
	this.TplName = "project/addlabel.html"


	projectId := this.GetIntParamFromURL(":ProjectId")
	projectLabelId := this.GetIntParamFromURL(":LabelId")
	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)

	projectLabelInfo, err := projectLabel.GetProjectLabelById(projectLabelId)
	this.checkError(err)


	this.Data["projectLabel"] = projectLabelInfo
	this.Data["projectInfo"] = projectInfo
	this.display()
}

// @router /project/dellabel/:ProjectId:int/:LabelId:int [get]
func (this *ProjectController)Dellabel()  {
	var projectLabel common.Project_label


	projectId := this.GetIntParamFromURL(":ProjectId")
	projectLabelId := this.GetIntParamFromURL(":LabelId")

	projectLabel.DeleteById(projectLabelId)

	this.redirect(beego.URLFor("ProjectController.Details",":Id", projectId))
}





// @router /project/ztree/:Id:int [get]
func (this *ProjectController)Ztree()  {
	this.Data["pageTitle"] = "项目树"
	var project common.Project
	this.TplName = "project/ztree.html"


	projectId := this.GetIntParamFromURL(":Id")
	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)

	ztree, err := libs.GetProjectZtreeByProjectId(projectId,projectInfo.Project_name)
	this.checkError(err)

	this.Data["ztree"] = ztree
}




