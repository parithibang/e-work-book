package controllers

import (
	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// MainController doc
type MainController struct {
	beego.Controller
}

// ActiveContent is the common template
func (c *MainController) ActiveContent(view string) {
	c.Layout = "basic-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Sidebar"] = "sidebar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = view + ".tpl"
	c.Data[view] = 1
}

// Get doc
func (c *MainController) Get() {
	c.ActiveContent("index")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}

// Search doc
func (c *MainController) Search() {
	c.ActiveContent("search")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}

// Employees doc
func (c *MainController) Employees() {
	c.ActiveContent("employees")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}

// Teams doc
func (c *MainController) Teams() {
	c.ActiveContent("teams")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}

// Pod doc
func (c *MainController) Pod() {
	c.ActiveContent("pod")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"

	o := orm.NewOrm()

	var pods []*models.Pods
	_, err := o.QueryTable(models.Pods{}).OrderBy("id").All(&pods)

	if err == nil {
		c.Data["POD"] = pods
	}
}

// Project doc
func (c *MainController) Project() {
	c.ActiveContent("project")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}

// Login doc
func (c *MainController) Login() {
	c.TplName = "login.tpl"
}

// LoginCheck doc
func (c *MainController) LoginCheck() {
	c.Redirect("/home", 302)
}

// Logout doc
func (c *MainController) Logout() {
	c.Redirect("/login", 302)
}

// Profile doc
func (c *MainController) Profile() {
	c.ActiveContent("profile")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}
