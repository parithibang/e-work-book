package controllers

import (
	"strconv"

	"github.com/parithibang/e-work-book/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
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

// UserSearchResults create a new user
func (c *SearchController) UserSearchResults() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	c.TplName = "search/search.tpl"
	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	userName := c.GetString("user-name")

	if userName == "" {
		flash.Set("custom_error", "User Name Should be given")
		flash.Store(&c.Controller)
		return
	}

	userProjects := models.UsersProjects{}
	searchList, count := userProjects.GetUserAssignedProjects(userName, pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)
	pageStart := currentPage*pageLimit - (pageLimit - 1)
	flash.Set("custom_redirect", beego.URLFor("SearchController.UserSearchResults", "user-name", userName))
	flash.Store(&c.Controller)

	c.Data["userName"] = userName
	c.Data["userProjectList"] = searchList
	c.Data["pageStart"] = pageStart
	c.Data["deleteMethod"] = "delete"
	c.Data["count"] = count
	c.Data["requestParams"] = beego.URLFor("SearchController.UserSearchResults", "user-name", userName)
}
