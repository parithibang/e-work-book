package controllers

import (
	"strconv"

	"github.com/parithibang/e-work-book/models"
	"github.com/parithibang/e-work-book/validations"
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

// Prepare user controller
func (c *UserController) Prepare() {
	c.BaseController.Prepare()
	c.Data["Title"] = "Users"
	c.Data["usersMenu"] = 1
}

// ListUsers to list all the users
func (c *UserController) ListUsers() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "users/list-users.tpl"

	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	users := models.Users{}
	userList, count := users.GetAllUsers(pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)
	pageStart := currentPage*pageLimit - (pageLimit - 1)

	c.Data["userList"] = userList
	c.Data["deleteMethod"] = "delete"
	c.Data["pageStart"] = pageStart
	c.Data["count"] = count
}

// AddUser to list user add form
func (c *UserController) AddUser() {
	pods := new(models.Pods).GetPods()
	c.Data["pods"] = pods
	c.Data["create"] = true
	c.Data["method"] = "post"
	c.TplName = "users/add-users.tpl"
}

// CreateUser create a new user
func (c *UserController) CreateUser() {
	o := orm.NewOrm()
	pods := new(models.Pods).GetPods()
	c.Data["pods"] = pods
	c.Data["create"] = true
	c.Data["method"] = "post"

	c.FormParsing()
	req := c.Ctx.Request
	valid, user := validations.Uservalidate(req)

	if !c.CheckErrors(valid, user) {
		flash := beego.NewFlash()
		_, err := o.Insert(&user)
		if err != nil {
			flash.Error("There is some problem while record insert")
		} else {
			flash.Success("User Created Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.TplName = "users/add-users.tpl"
}

// EditUser to list user add form
func (c *UserController) EditUser() {
	flash := beego.NewFlash()
	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	user := models.Users{Id: userId}
	err := o.Read(&user)

	if err != nil {
		flash.Error("User Not Found!!!")
		flash.Store(&c.Controller)
		c.Abort("404")
	}

	pods := new(models.Pods).GetPods()
	c.Data["User"] = user
	c.Data["pods"] = pods
	c.Data["update"] = true
	c.Data["method"] = "put"
	c.TplName = "users/add-users.tpl"
}

// UpdateUser to list user add form
func (c *UserController) UpdateUser() {
	o := orm.NewOrm()
	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	user := models.Users{Id: userId}
	o.Read(&user)
	pods := new(models.Pods).GetPods()
	c.FormParsing()
	req := c.Ctx.Request
	valid, updatedUser := validations.Uservalidate(req)
	updatedUser.Id = userId

	if !c.CheckErrors(valid, updatedUser) {
		flash := beego.NewFlash()
		_, err := o.Update(&updatedUser)
		if err != nil {
			flash.Error("There is some problem while record update")
		} else {
			flash.Success("User Updated Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.Data["User"] = updatedUser
	c.Data["pods"] = pods
	c.Data["update"] = true
	c.Data["method"] = "put"

	c.TplName = "users/add-users.tpl"
}

//FormParsing checks form parsing
func (c *UserController) FormParsing() {
	user := models.Users{}
	if err := c.ParseForm(&user); err != nil {
		flash := beego.NewFlash()
		flash.Error("Form Parsing Failed")
		flash.Store(&c.Controller)

		return
	}
}

//CheckErrors checks errors while validating
func (c *UserController) CheckErrors(valid validation.Validation, user models.Users) bool {
	if valid.HasErrors() {
		flash := beego.NewFlash()
		flash.Error("Validation Failed!!")
		flash.Store(&c.Controller)

		c.Data["Errors"] = valid.ErrorsMap
		c.Data["User"] = user
		return true
	}

	return false
}

// DeleteUser to delete user
func (c *UserController) DeleteUser() {
	o := orm.NewOrm()
	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	flash := beego.NewFlash()

	_, err := o.QueryTable(models.Users{}).Filter("id", userId).Update(orm.Params{
		"is_active": 0,
	})

	if err != nil {
		flash.Set("custom_error", "There is some problem while deleteing record")
	} else {
		flash.Set("custom_success", "User Deleted Successfully!!!")
	}

	flash.Store(&c.Controller)

	c.Redirect(beego.URLFor("UserController.ListUsers"), 303)
}
