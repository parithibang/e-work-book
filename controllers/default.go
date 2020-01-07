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
	c.Layout = "basic-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Sidebar"] = "sidebar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = view + ".tpl"
}

// Get doc
func (c *MainController) Get() {
	c.ActiveContent("index")
	c.Data["Website"] = "eWorkBook"
	c.Data["Email"] = "eworkbook@gmail.com"
}
