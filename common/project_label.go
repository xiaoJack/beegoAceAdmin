package common

import (
	"github.com/astaxie/beego/validation"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

type Project_label struct {
	Id         int
	Project_id         int   `orm:"size(11)"      form:"Project_id"       valid:"Required"    validmsg:"项目ID不能为空"`
	Label_name   	   string    `orm:"size(20)"  form:"Label_name"       valid:"Required;MaxSize(20)"    validmsg:"标签名称不能为空最长不能超过20个字符"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`   // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`       // 更新时间
}

func (this *Project_label) Valid()(b bool, str string)  {
	v := validation.Validation{}
	b, _ = v.Valid(this)
	if !b{

		st := reflect.TypeOf(Project_label{})
		field,_ := st.FieldByName(v.Errors[0].Field)
		msg := field.Tag.Get("validmsg")

		return b, msg
	}
	return true, ""
}

// 根据ID获取项目配置信息
func (this *Project_label) GetProjectLabelById(Id int) (*Project_label, error) {
	p := &Project_label{}
	p.Id = Id
	err := o.Read(p)
	return p, err
}


// 添加
func (this *Project_label) Add() (*Project_label, error) {

	_, err := o.Insert(this)
	return this, err
}



//删除
func (this *Project_label)DeleteById(Id int)(int64, error)  {
	p := &Project_label{}
	p.Id = Id
	del, err := o.Delete(p)
	return del, err
}




// 根据ID获取项目配置信息
func (this *Project_label) GetListByProjectId(ProjectId int) ([]Project_label, error) {

	var list []Project_label
	//err := o.Read(list, "project_id")


	qs := o.QueryTable(TableName("project_label"))
	_, err := qs.Filter("project_id",ProjectId).OrderBy("id").All(&list)



	return list, err
}



// 更新用户信息
func (this *Project_label) UpdateProject(fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New("更新字段不能为空")
	}
	_, err := o.Update(this, fileds...)
	return err
}