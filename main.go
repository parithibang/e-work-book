package main

import (
	_ "github.com/aravindkumaremis/e-work-book/routers"
	_ "github.com/aravindkumaremis/e-work-book/setup"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

// incrementByOne function will increment by one to the given integer
func incrementByOne(in int) (out int) {
	out = in + 1
	return
}

func main() {
	beego.AddFuncMap("incrementByOne", incrementByOne)
	beego.Run()
}
