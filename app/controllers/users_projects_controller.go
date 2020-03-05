package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
	"github.com/parithibang/e-work-book/app/validations"
)

// UsersProjectsController Controller actions for users assigned projects
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

		userPercentage := new(models.UsersProjects).GetTotalWorkPercentageOfUser(userProjects.Users.ID)
		percentage := fmt.Sprintf("%v", userPercentage[0])
		percentageFloat, _ := strconv.ParseFloat(percentage, 64)
		displayUser := models.Users{ID: userProjects.Users.ID}
		_ = o.Read(&displayUser)

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

// EditUserProjects edit the user project record
func (c *UsersProjectsController) EditUserProjects() {
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data["custom_redirect"])
	redirectURL := flash.Data["custom_redirect"]
	userProjectID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	userProject := models.UsersProjects{ID: userProjectID}
	projects := new(models.Projects).GetProjects()
	users := new(models.Users).GetUsers()

	err := o.Read(&userProject)

	if err != nil {
		flash.Error("UsersProject Not Found!!!")
		flash.Store(&c.Controller)
		c.Abort("404")
	}

	c.Data["projects"] = projects
	c.Data["users"] = users
	c.Data["UserProjects"] = userProject
	c.Data["update"] = true
	c.Data["method"] = "put"
	c.Data["redirectURL"] = redirectURL
	c.TplName = "users-projects/add-users-projects.tpl"
}

// UpdateUserProjects to update the user project record
func (c *UsersProjectsController) UpdateUserProjects() {
	o := orm.NewOrm()
	redirectLink := c.GetString("redirect-link")
	fmt.Println(redirectLink)
	projects := new(models.Projects).GetProjects()
	users := new(models.Users).GetUsers()
	userProjectID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	userProject := models.UsersProjects{ID: userProjectID}
	_ = o.Read(&userProject)
	req := c.Ctx.Request
	valid, updatedUserProject := validations.UserProjectvalidate(req)
	updatedUserProject.ID = userProjectID

	if !c.CheckErrors(valid, updatedUserProject) {
		flash := beego.NewFlash()
		_, err := o.Update(&updatedUserProject)

		userPercentage := new(models.UsersProjects).GetTotalWorkPercentageOfUser(updatedUserProject.Users.ID)
		percentage := fmt.Sprintf("%v", userPercentage[0])
		percentageFloat, _ := strconv.ParseFloat(percentage, 64)
		displayUser := models.Users{ID: updatedUserProject.Users.ID}
		_ = o.Read(&displayUser)

		if percentageFloat > 100.00 {
			flash.Warning("%s is allocated with %.2f%s of work", displayUser.FullName(), percentageFloat, "%")
		}

		if err != nil {
			flash.Error("There is some problem while record update")
		} else {
			flash.Success("Record Updated Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.Data["UserProjects"] = updatedUserProject
	c.Data["projects"] = projects
	c.Data["users"] = users
	c.Data["update"] = true
	c.Data["method"] = "put"
	c.Data["redirectURL"] = redirectLink

	c.TplName = "users-projects/add-users-projects.tpl"
}

// DeleteUserProject to delete user
func (c *UsersProjectsController) DeleteUserProject() {
	o := orm.NewOrm()
	userProjectID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	flash := beego.NewFlash()
	redirectURL := c.GetString("request-params")
	_, err := o.QueryTable(models.UsersProjects{}).Filter("id", userProjectID).Update(orm.Params{
		"is_active": 0,
	})

	if err != nil {
		flash.Set("custom_error", "There is some problem while deleteing record")
	} else {
		flash.Set("custom_success", "Record Deleted Successfully!!!")
	}

	flash.Store(&c.Controller)

	c.Redirect(redirectURL, 303)
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
