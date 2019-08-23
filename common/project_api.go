package common

import (
	"github.com/astaxie/beego/validation"
	"github.com/pkg/errors"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
)

type Project_api struct {
	Id                 int        `orm:"auto"`
	Project_id         int        `orm:"size(11)"      form:"Project_id"             valid:"Required"                validmsg:"项目ID不能为空"`
	Project_label_id   int        `orm:"size(11)"      form:"Project_label_id"       valid:"Required"                validmsg:"标签ID不能为空"`
	Api_name   	       string     `orm:"size(20)"      form:"Api_name"               valid:"Required;MaxSize(20)"    validmsg:"接口名称不能为空且长度不能超过20个字符"`
	Api_url            string     `orm:"size(50)"      form:"Api_url"                valid:"Required;MaxSize(50)"    validmsg:"接口URL不能为空且长度不能超过50个字符"`
	Method             int        `orm:"size(2)"       form:"Method"                 valid:"Required"                validmsg:"请求方法不能为空"`
	Intro              string     `orm:"size(255)"     form:"Intro" `
	Status             int        `orm:"size(2)"       form:"Status"`
	CreateTime         time.Time  `orm:"auto_now_add;type(datetime)"`   // 创建时间
	UpdateTime         time.Time  `orm:"auto_now;type(datetime)"`       // 更新时间
}


func (this *Project_api) Valid()(b bool, str string)  {
	v := validation.Validation{}
	b, _ = v.Valid(this)
	if !b{

		st := reflect.TypeOf(Project_api{})
		field,_ := st.FieldByName(v.Errors[0].Field)
		msg := field.Tag.Get("validmsg")

		return b, msg
	}
	return true, ""
}



// 添加
func (this *Project_api) Add() (*Project_api, error) {

	_, err := o.Insert(this)
	return this, err
}


// 更新
func (this *Project_api) UpdateProject(fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New("更新字段不能为空")
	}
	_, err := o.Update(this, fileds...)
	return err
}



// 根据ID获取信息
func (this *Project_api) GetProjectApiById(Id int) (*Project_api, error) {
	p := &Project_api{}
	p.Id = Id
	err := o.Read(p)
	return p, err
}

func (this *Project_api)GetListByProjectId(projectId int)(list []Project_api, err error)  {
	o := orm.NewOrm()

	 _,err = o.QueryTable(TableName("project_api")).Filter("Project_id", projectId).OrderBy("id").All(&list)

	return list, err
}


