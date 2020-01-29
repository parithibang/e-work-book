package controllers

import (
	"github.com/astaxie/beego"
)

// BaseController doc
type BaseController struct {
	beego.Controller
}

// Prepare to prepare all common entities
func (c *BaseController) Prepare() {
	c.CommonViews()
}

// CommonViews to set up the common views
func (c *BaseController) CommonViews() {
	c.Layout = "base/basic-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "base/header.tpl"
	c.LayoutSections["Sidebar"] = "base/sidebar.tpl"
	c.LayoutSections["Footer"] = "base/footer.tpl"
}
