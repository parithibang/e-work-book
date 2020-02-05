package controllers

import (
	"strconv"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego/validation"
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
	c.Data["users"] = 1
}

// AddUser to list user add form
func (c *UserController) AddUser() {
	o := orm.NewOrm()
	var pods []*models.Pods
	o.QueryTable(models.Pods{}).OrderBy("name").All(&pods, "id", "name")
	c.Data["pods"] = pods
	c.TplName = "users/add-users.tpl"
}

// CreateUser create a new user
func (c *UserController) CreateUser() {
	o := orm.NewOrm()
	flash := beego.NewFlash()
	var pods []*models.Pods
	o.QueryTable(models.Pods{}).OrderBy("name").All(&pods, "id", "name")
	c.Data["pods"] = pods

	podId, _ := strconv.Atoi(c.GetString("pods"))

	selectedPod := models.Pods{
		Id: podId,
	}

	user := models.Users{
		FirstName: c.GetString("first-name"),
		LastName:  c.GetString("last-name"),
		Password:  c.GetString("password"),
		UserName:  c.GetString("user-name"),
		Pods:      &selectedPod,
	}

	valid := validation.Validation{}
	valid.Valid(&user)

	if podId == 0 {
		valid.SetError("Pods", "Pod Should be selected")
	}

	if valid.HasErrors() {

		flash.Error("Validation Failed!!")
		flash.Store(&c.Controller)

		c.Data["Errors"] = valid.ErrorsMap
		c.Data["User"] = user
	} else {
		_, err := o.Insert(&user)
		if err != nil {
			flash.Error("There is some problem while record insert")
			flash.Store(&c.Controller)
		} else {
			flash.Success("User Created Successfully!!!")
			flash.Store(&c.Controller)
		}
	}

	c.TplName = "users/add-users.tpl"
}
