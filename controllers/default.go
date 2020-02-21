package controllers

import (
	"github.com/astaxie/beego"
)

// MainController doc
type MainController struct {
	beego.Controller
}

// ActiveContent is the common template
func (c *MainController) ActiveContent(view string) {
	c.Layout = "base/basic-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "base/header.tpl"
	c.LayoutSections["Sidebar"] = "base/sidebar.tpl"
	c.LayoutSections["Footer"] = "base/footer.tpl"
	c.TplName = view + ".tpl"
	c.Data[view] = 1
}

// Get doc
func (c *MainController) Get() {
	c.ActiveContent("index")
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
