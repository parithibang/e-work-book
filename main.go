package main

import (
	_ "github.com/aravindkumaremis/e-work-book/routers"
	_ "github.com/aravindkumaremis/e-work-book/setup"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}
