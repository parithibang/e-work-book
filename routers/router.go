package routers

import (
	"github.com/aravindkumaremis/e-work-book/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{}, "get,post:Get")
	beego.Router("/search", &controllers.MainController{}, "get,post:Search")
	beego.Router("/employees", &controllers.MainController{}, "get,post:Employees")
	beego.Router("/teams", &controllers.MainController{}, "get,post:Teams")
	beego.Router("/pod", &controllers.MainController{}, "get,post:Pod")
	beego.Router("/project", &controllers.MainController{}, "get,post:Project")
	beego.Router("/login", &controllers.MainController{}, "get,post:Login")
	beego.Router("/loginCheck", &controllers.MainController{}, "get,post:LoginCheck")
	beego.Router("/logout", &controllers.MainController{}, "get,post:Logout")
	beego.Router("/profile", &controllers.MainController{}, "get,post:Profile")
}
