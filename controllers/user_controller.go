package controllers

import (
	"strconv"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego/utils/pagination"
)

// UserController doc
type UserController struct {
	BaseController
}

const pageLimit = 10

// ListUsers to list all the users
func (c *UserController) ListUsers() {
	c.TplName = "users/list-users.tpl"

	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	users := models.Users{}
	userList, count := users.GetAllUsers(pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)

	c.Data["userList"] = userList
}
