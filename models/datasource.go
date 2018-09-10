package models

import (
	"fmt"

	"github.com/astaxie/beego"
)

type datasource struct {
	name   string
	host   string
	port   int
	user   string
	pwd    string
	driver string
}

func lucky_drawDbName() string {
	return beego.AppConfig.String("db")
}

func lucky_drawDataSource() (string, string) {
	ds := datasource{
		name:   beego.AppConfig.String("db"),
		host:   beego.AppConfig.String("host"),
		port:   beego.AppConfig.DefaultInt("port", 3306),
		user:   beego.AppConfig.String("user"),
		pwd:    beego.AppConfig.String("pwd"),
		driver: beego.AppConfig.String("driver"),
	}
	return ds.driver, ds.ConnString()
}

// ConnString 根据驱动名返回不同的数据库链接字符串
func (ds *datasource) ConnString() (connString string) {
	switch ds.driver {
	case "mssql":
		connString = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s", ds.host, ds.name, ds.user, ds.pwd)
		if ds.port > 0 {
			connString = fmt.Sprintf("%s;port=%d", connString, ds.port)
		}
		connString = connString + ";encrypt=disable"
	case "mysql":
		connString = fmt.Sprintf("%s:%s@tcp(%s", ds.user, ds.pwd, ds.host)
		if ds.port > 0 {
			connString = fmt.Sprintf("%s:%d)/%s", connString, ds.port, ds.name)
		} else {
			connString = fmt.Sprintf("%s)/%s", connString, ds.name)
		}
		connString = fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", connString)
	case "postgres":
		connString = fmt.Sprintf("host=%s ", ds.host)
		if ds.port > 0 {
			connString = fmt.Sprintf("%s:%d", connString, ds.port)
		}
		connString = fmt.Sprintf("%s user=%s dbname=%s sslmode=disable password=%s", connString, ds.user, ds.name, ds.pwd)
	}
	return
}
