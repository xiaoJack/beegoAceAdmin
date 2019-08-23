package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

var (
	o                 orm.Ormer
	tablePrefix       string             // 表前缀
	UserService       *User       // 用户服务
)

func Init() {
	dbHost := beego.AppConfig.String("DBHost")
	dbPort := beego.AppConfig.String("DBPort")
	dbUser := beego.AppConfig.String("DBUser")
	dbPassword := beego.AppConfig.String("DBPasswd")
	dbName := beego.AppConfig.String("DBName")
	timezone := beego.AppConfig.String("DBTimezone")
	tablePrefix = beego.AppConfig.String("DBPrefix")
	dbmaxIdle := beego.AppConfig.DefaultInt("DBMaxIdle", 30)
	dbmaxConn := beego.AppConfig.DefaultInt("DBMaxConn", 30)


	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn, dbmaxIdle, dbmaxConn)

	orm.RegisterModelWithPrefix(tablePrefix,
		new(User),
		new(Project),
		new(Project_label),
		new(Project_api),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	initService()

}


func initService() {
	UserService = &User{}
}



// 返回真实表名
func TableName(name string) string {
	return tablePrefix + name
}


func DBVersion() string {
	var lists []orm.ParamsList
	o.Raw("SELECT VERSION()").ValuesList(&lists)
	return lists[0][0].(string)
}
