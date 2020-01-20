package setup

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func db() {
	dbAlias := beego.AppConfig.String("db_alias")
	dbUserName := beego.AppConfig.String("db_user_name")
	dbPassword := beego.AppConfig.String("db_password")
	dbPort := beego.AppConfig.String("db_port")
	dbHost := beego.AppConfig.String("db_host")
	dbCharset := beego.AppConfig.String("db_charset")
	dbName := beego.AppConfig.String("db_name")

	orm.RegisterDataBase(dbAlias, "mysql", dbUserName+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)

	isDev := (beego.AppConfig.String("runmode") == "dev")

	//debug querys
	if isDev {
		orm.Debug = isDev
	}
}
