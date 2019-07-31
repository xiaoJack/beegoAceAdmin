package common

import (
	"github.com/astaxie/beego/utils"
	"github.com/lisijie/gopub/app/libs"
	"github.com/pkg/errors"
	"time"
)

type User struct {
	Id         int
	UserName   string    `orm:"unique;size(20)"`             // 用户名
	Password   string    `orm:"size(32)"`                    // 密码
	Salt       string    `orm:"size(10)"`                    // 密码盐
	Sex        int       `orm:"default(0)"`                  // 性别
	Email      string    `orm:"size(50)"`                    // 邮箱
	LastLogin  time.Time `orm:"null;type(datetime)"`         // 最后登录时间
	LastIp     string    `orm:"size(15)"`                    // 最后登录IP
	Status     int       `orm:"default(0)"`                  // 状态，0正常 -1禁用
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`     // 更新时间
//	RoleList   []Role    `orm:"-"`                           // 角色列表
}



// 分页获取用户列表
func (this *User) GetUserList(page, pageSize int) ([]User, error) {
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	var users []User
	qs := o.QueryTable(TableName("user"))
	_, err := qs.OrderBy("id").Limit(pageSize, offset).All(&users)

	return users, err
}



// 添加用户
func (this *User) AddUser(userName, email, password string, sex int, status int) (*User, error) {
	if exists, _ := this.GetUserByName(userName); exists.Id > 0 {
		return nil, errors.New("用户名已存在")
	}

	user := &User{}
	user.UserName = userName
	user.Sex = sex
	user.Email = email
	user.Status = status
	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = libs.Md5([]byte(password + user.Salt))
	// user.LastLogin = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	_, err := o.Insert(user)
	return user, err
}



// 获取用户总数
func (this *User) GetTotal() (int64, error) {
	return o.QueryTable(TableName("user")).Count()
}


// 根据用户名获取用户信息
func (this *User) GetUserByName(userName string) (*User, error) {
	user := &User{}
	user.UserName = userName
	err := o.Read(user, "UserName")
	return user, err
}


// 根据用户名获取用户信息
func (this *User) GetUserById(Id int) (*User, error) {
	user := &User{}
	user.Id = Id
	err := o.Read(user)
	return user, err
}


//删除用户
func (this *User)DeleteById(Id int)(int64, error)  {
	user := &User{}
	user.Id = Id
	del, err := o.Delete(user)
	return del, err
}



// 更新用户信息
func (this *User) UpdateUser(user *User, fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New("更新字段不能为空")
	}
	_, err := o.Update(user, fileds...)
	return err
}