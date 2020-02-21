package controllers

import (
	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego"
)

// SearchController doc
type SearchController struct {
	BaseController
}

// Prepare teams controller
func (c *SearchController) Prepare() {
	c.BaseController.Prepare()
	c.Data["Title"] = "Search"
	c.Data["searchMenu"] = 1
}

// GetSearch to list user add formx
func (c *SearchController) GetSearch() {
	c.TplName = "search/search.tpl"
}

// PostSearchResults create a new user
func (c *SearchController) PostSearchResults() {
	c.TplName = "search/search.tpl"

	userName := c.GetString("user-name")

	if userName == "" {
		flash := beego.NewFlash()
		flash.Error("User Name Should be given")
		flash.Store(&c.Controller)
		return
	}

	userProjects := models.UsersProjects{}
	searchList := userProjects.GetUserAssignedProjects(userName)

	c.Data["userName"] = userName
	c.Data["userList"] = searchList
}
