package controllers

import (
	"admin/common"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"time"
)

type ProjectController struct {
	BaseController
}

var project common.Project


// @router /project/list [get]
func (this *ProjectController)List()  {

	this.Data["pageTitle"] = "API项目管理"


	list, _ := project.GetprojectList()

	this.Data["list"] = list
	this.display()
}



// @router /project/add [get,post]
func (this *ProjectController)Add()  {

	this.Data["pageTitle"] = "添加项目"
	if this.isPost() {
		valid := validation.Validation{}

		//project.Project_name = this.GetString("project_name")
		//project.Project_describe = this.GetString("project_describe")
		//project.Project_url = this.GetString("project_url")
		//project.Test_ip = this.GetString("test_ip")
		//project.Release_ip = this.GetString("release_ip")
		//project.Pro_ip = this.GetString("pro_ip")
		//
		//is_monitor,_ := this.GetInt8("is_monitor", 1)
		//project.Is_monitor = is_monitor
		//
		//project.Monitor_url = this.GetString("monitor_url")

		this.ParseForm(&project)
		fmt.Println(project)
		//fmt.Printf("valid ---start---\n")
		//b, err := project.Valid()
		//if !b{
		//	fmt.Println(err)
		//	this.showMsg(err , MSG_ERR)
		//	//for _, e := range err.Errors {
		//	//	this.showMsg(e.Message, MSG_ERR)
		//	//}
		//}
		//fmt.Printf("valid ---end---\n")

		valid.Required(project.Project_name, "project_name").Message("请输入项目名称")
		valid.Required(project.Project_url, "project_url").Message("请输入项目域名")

		valid.IP(project.Test_ip, "test_ip").Message("请输入正确的IP地址")
		valid.IP(project.Release_ip, "release_ip").Message("请输入正确的IP地址")
		valid.IP(project.Pro_ip, "pro_ip").Message("请输入正确的IP地址")


		if valid.HasErrors() {
			for _, err := range valid.Errors {
				fmt.Println(err)
				this.showMsg(err.Message, MSG_ERR)
			}
		}


		Id, _ := this.GetInt("Id",0)
		if Id == 0{
			//insert
			project.CreateTime = time.Now()
			project.UpdateTime = time.Now()
			_, err := project.Add()
			this.checkError(err)
		}else{
			//update
			project.Id = Id
			project.UpdateTime = time.Now()
			err := project.UpdateProject("Project_name", "Project_describe", "Project_url", "Test_ip", "Release_ip", "Pro_ip", "Is_monitor", "Monitor_url")
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