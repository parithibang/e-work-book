package controllers

import (
	"strconv"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

// SearchController doc
type SearchController struct {
	BaseController
}

// GetSearch to list user add form
func (c *SearchController) GetSearch() {
	c.Data["searchmenu"] = 1
	c.TplName = "search/search.tpl"
}

// PostSearchResults create a new user
func (c *SearchController) PostSearchResults() {
	c.TplName = "search/search.tpl"
	c.Data["searchmenu"] = 1

	userName := c.GetString("user-name")

	if userName == "" {
		flash := beego.NewFlash()
		flash.Error("User Name Should be given")
		flash.Store(&c.Controller)
		return
	}

	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	userProjects := models.UsersProjects{}
	userList, count := userProjects.GetUserAssignedProjects(userName, pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)

	c.Data["userName"] = userName
	c.Data["userList"] = userList
}
