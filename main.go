package main

import (
	_ "github.com/parithibang/e-work-book/app/routers"
	_ "github.com/parithibang/e-work-book/setup"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}
