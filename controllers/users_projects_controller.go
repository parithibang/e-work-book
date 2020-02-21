package controllers

import (
	"fmt"
	"strconv"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/aravindkumaremis/e-work-book/validations"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// SearchController doc
type UsersProjectsController struct {
	BaseController
}

// Prepare teams controller
func (c *UsersProjectsController) Prepare() {
	c.BaseController.Prepare()
	c.Data["Title"] = "User Projects"
	c.Data["searchMenu"] = 1
}

// AddUserProjectDetail add user project
func (c *UsersProjectsController) AddUserProjectDetail() {
	// userPercentage := new(models.UsersProjects).GetTotalWorkPercentageOfUser(2)

	// str := fmt.Sprintf("%v", userPercentage[0])
	// fmt.Println(str)

	users := new(models.Users).GetUsers()
	projects := new(models.Projects).GetProjects()

	c.Data["projects"] = projects
	c.Data["users"] = users
	c.Data["create"] = true
	c.Data["method"] = "post"
	c.TplName = "users-projects/add-users-projects.tpl"
}

// CreateUserProjectDetail add user project
func (c *UsersProjectsController) CreateUserProjectDetail() {
	o := orm.NewOrm()
	projects := new(models.Projects).GetProjects()
	users := new(models.Users).GetUsers()
	c.Data["projects"] = projects
	c.Data["users"] = users
	c.Data["create"] = true
	c.Data["method"] = "post"

	req := c.Ctx.Request
	valid, userProjects := validations.UserProjectvalidate(req)

	if !c.CheckErrors(valid, userProjects) {
		flash := beego.NewFlash()
		_, err := o.Insert(&userProjects)

		userPercentage := new(models.UsersProjects).GetTotalWorkPercentageOfUser(userProjects.Users.Id)
		percentage := fmt.Sprintf("%v", userPercentage[0])
		percentageFloat, _ := strconv.ParseFloat(percentage, 64)
		displayUser := models.Users{Id: userProjects.Users.Id}
		o.Read(&displayUser)

		if percentageFloat > 100.00 {
			flash.Warning("%s is allocated with %.2f%s of work", displayUser.FullName(), percentageFloat, "%")
		}

		if err != nil {
			flash.Error("There is some problem while record insert")
		} else {
			flash.Success("User Project Added Successfully!!!")
		}

		flash.Store(&c.Controller)
	}

	c.TplName = "users-projects/add-users-projects.tpl"
}

//CheckErrors checks errors while validating
func (c *UsersProjectsController) CheckErrors(valid validation.Validation, userprojects models.UsersProjects) bool {
	if valid.HasErrors() {
		flash := beego.NewFlash()
		flash.Error("Validation Failed!!")
		flash.Store(&c.Controller)

		c.Data["Errors"] = valid.ErrorsMap
		c.Data["UserProjects"] = userprojects
		return true
	}

	return false
}
