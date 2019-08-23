package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xiaoJack/beegoAceAdmin/common"
)

type ApiController struct {
	BaseController
}


var ApiMethod = map[int]string{
	1:"GET",
	2:"POST",
	//3:"PUT",
	//4:"PATCH",
	//5:"DELETE",
}

var ApiStatus = map[int]string{
	0:"开发中",
	1:"上线",
	2:"停止使用",
}

// @router /api/add/:ProjectId:int/:LabelId:int [get,post]
func (this *ApiController)Add()  {

	this.Data["pageTitle"] = "添加接口"

	var project common.Project
	var projectLabel common.Project_label
	var projectApi common.Project_api

	projectId := this.GetIntParamFromURL(":ProjectId")
	projectLabelId := this.GetIntParamFromURL(":LabelId")
	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)

	projectLabelInfo, err := projectLabel.GetProjectLabelById(projectLabelId)
	this.checkError(err)


	if this.isPost() {

		this.ParseForm(&projectApi)
		b, err := projectApi.Valid()
		if !b{
			this.Data["projectLabelInfo"] = projectLabelInfo
			this.Data["projectInfo"] = projectInfo
			this.Data["ApiMethod"] = ApiMethod
			this.Data["ApiStatus"] = ApiStatus
			this.showMsg(err , MSG_ERR)
		}


		Id, _ := this.GetInt("Id",0)
		if Id == 0{
			//insert
			_, err := projectApi.Add()
			this.checkError(err)
		}else{
			//update
			projectApi.Id = Id
			err := projectApi.UpdateProject("Api_name", "Api_url", "Method", "Intro", "Status")
			this.checkError(err)
		}



		this.redirect(beego.URLFor("ProjectController.Details",":Id", projectId))
	}

	this.Data["projectLabelInfo"] = projectLabelInfo
	this.Data["projectInfo"] = projectInfo
	this.Data["ApiMethod"] = ApiMethod
	this.Data["ApiStatus"] = ApiStatus
	this.display()
}



// @router /api/edit/:ProjectId:int/:LabelId:int/:ApiId:int [get]
func (this *ApiController)Edit()  {
	this.Data["pageTitle"] = "编辑接口"
	this.TplName = "api/add.html"
	var project common.Project
	var projectLabel common.Project_label
	var projectApi common.Project_api

	projectId := this.GetIntParamFromURL(":ProjectId")
	projectLabelId := this.GetIntParamFromURL(":LabelId")
	apiId := this.GetIntParamFromURL(":ApiId")


	projectLabelInfo, err := projectLabel.GetProjectLabelById(projectLabelId)
	this.checkError(err)

	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)

	apiInfo, err := projectApi.GetProjectApiById(apiId)
	this.checkError(err)


	fmt.Println(projectInfo)
	this.Data["projectInfo"] = projectInfo
	this.Data["projectLabelInfo"] = projectLabelInfo
	this.Data["apiInfo"] = apiInfo
	this.Data["ApiMethod"] = ApiMethod
	this.Data["ApiStatus"] = ApiStatus
	this.display()

}


// @router /api/details/:ProjectId:int/:LabelId:int [get]
func (this *ApiController)Details()  {
	this.Data["pageTitle"] = "接口详情"
	var project common.Project


	projectId := this.GetIntParamFromURL(":Id")
	projectInfo, err := project.GetProjectById(projectId)
	this.checkError(err)


	this.Data["projectInfo"] = projectInfo

	this.display()
}



