package common

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/pkg/errors"
	"time"
)

type Project struct {
	Id         int
	Project_name   	   string    `orm:"size(45)"  form:"project_name"       valid:"Required"`	// 项目名称,必输项
	Project_describe   string    `orm:"size(100)" form:"project_describe"`             			// 项目描述
	Project_url        string    `orm:"size(100)" form:"project_url"		valid:"Required"`   // 项目域名,必输项
	Test_ip            string    `orm:"size(15)"  form:"test_ip"			valid:"IP"`         // 测试环境IP,IP格式
	Release_ip         string    `orm:"size(15)"  form:"release_ip"			valid:"IP"`         // 预发布环境IP,IP格式
	Pro_ip             string    `orm:"size(15)"  form:"pro_ip"				valid:"IP"`         // 生产环境IP,IP格式
	Is_monitor         int8      `orm:"size(2)"   form:"is_monitor"`               				// 1开启监控，默认开启
	Monitor_url        string    `orm:"size(15)"  form:"monitor_url"`              				// 监控触发URL
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`   // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`       // 更新时间
}

func (this *Project) Valid()(b bool, str string)  {
	v := validation.Validation{}
	b, _ = v.Valid(this)
	if !b{
		fmt.Println(v.Errors)
		return b, v.Errors[0].Key + v.Errors[0].Message
	}
	return true, ""
}




// 获取所有的项目列表
func (this *Project) GetprojectList() ([]Project, error) {

	var list []Project
	qs := o.QueryTable(TableName("project"))
	_, err := qs.OrderBy("id").All(&list)

	return list, err
}


// 添加用户
func (this *Project) Add() (*Project, error) {

	_, err := o.Insert(this)
	return this, err
}


// 根据ID获取项目配置信息
func (this *Project) GetProjectById(Id int) (*Project, error) {
	p := &Project{}
	p.Id = Id
	err := o.Read(p)
	return p, err
}



// 更新用户信息
func (this *Project) UpdateProject(fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New("更新字段不能为空")
	}
	_, err := o.Update(this, fileds...)
	return err
}