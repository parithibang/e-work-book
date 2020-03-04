package setup

import "github.com/astaxie/beego"

func init() {
	beego.SetStaticPath("/static", "app/static")
	db()
	setDefaultMessage()
	functionsForView()
	filters()
}
