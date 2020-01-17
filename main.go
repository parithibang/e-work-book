package main

import (
	_ "e-work-book/routers"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbAlias       = "default"
	mysqlUser     = "root"
	mysqlPassword = "rootroot"
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlDatabase = "testDB"
	mysqlCharset  = "utf8"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)
)

func init() {
	// register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	err := orm.RegisterDataBase(dbAlias, "mysql", mysqlCon)
	if err != nil {
		log.Println("Database connection failed!")
		panic(err)
	} else {
		fmt.Println("Database connected successfully!")
	}
}
func main() {
	beego.Run()
}
