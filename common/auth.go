package common

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// 登录验证服务
type AuthService struct {
	loginUser *User    // 当前登录用户
//	permMap   map[string]bool // 当前用户权限表
//	openPerm  map[string]bool // 公开的权限
}

func NewAuth() *AuthService {
	return new(AuthService)
}

// 初始化开放权限
func (this *AuthService) initOpenPerm() {
	//this.openPerm = map[string]bool{
	//	"main.index":        true,
	//	"main.profile":      true,
	//	"main.login":        true,
	//	"main.logout":       true,
	//	"main.getpubstat":   true,
	//	"project.clone":     true,
	//	"project.getstatus": true,
	//	"task.gettags":      true,
	//	"task.getstatus":    true,
	//	"task.startpub":     true,
	//}
}




// 初始化
func (this *AuthService) Init(token string) {
	//this.initOpenPerm()
	//arr := strings.Split(token, "|")
	//beego.Trace("登录验证, token: ", token)
	//if len(arr) == 2 {
	//	idstr, password := arr[0], arr[1]
	//	userId, _ := strconv.Atoi(idstr)
	//	if userId > 0 {
	//		user, err := UserService.GetUser(userId, true)
	//		if err == nil && password == libs.Md5([]byte(user.Password+user.Salt)) {
	//			this.loginUser = user
	//			this.initPermMap()
	//			beego.Trace("验证成功，用户信息: ", user)
	//		}
	//	}
	//}
}


// 生成md5
func getMd5Str(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//密码MD5值
func GetPassowrdMd5Str(str string)string  {
	return getMd5Str([]byte(str))
}



// 用户登录
func (this *AuthService) Login(userName, password string) (*User, error) {

	user, err := UserService.GetUserByName(userName)

	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("帐号或密码错误")
		} else {
			return nil, errors.New("系统错误")
		}
	}

	if user.Password != GetPassowrdMd5Str(password) {
		return nil, errors.New("帐号或密码错误2")
	}
	if user.Status == -1 {
		return nil, errors.New("该帐号已禁用")
	}

	user.LastLogin = time.Now()
	UserService.UpdateUser(user, "LastLogin")

	return user, nil
}

// 退出登录
func (this *AuthService) Logout() error {
	return nil
}
