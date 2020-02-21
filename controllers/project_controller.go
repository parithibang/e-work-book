package controllers

import (
	"strconv"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/aravindkumaremis/e-work-book/validations"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego/validation"
)

// ProjectController doc
type ProjectController struct {
	BaseController
}

// Prepare teams controller
func (c *ProjectController) Prepare() {
	c.BaseController.Prepare()
	c.Data["Title"] = "Projects"
	c.Data["projectMenu"] = 1
}

// ListProjects to list all the projects
func (c *ProjectController) ListProjects() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "projects/list-projects.tpl"

	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	projects := models.Projects{}
	projectList, count := projects.GetAllProjects(pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)
	pageStart := currentPage*pageLimit - (pageLimit - 1)

	c.Data["projectList"] = projectList
	c.Data["deleteMethod"] = "delete"
	c.Data["pageStart"] = pageStart
	c.Data["count"] = count
}

// AddProject to list team add form
func (c *ProjectController) AddProject() {
	c.Data["create"] = true
	c.Data["method"] = "post"
	c.TplName = "projects/add-projects.tpl"
}

// CreateProject to list team add form
func (c *ProjectController) CreateProject() {
	o := orm.NewOrm()
	c.Data["create"] = true
	c.Data["method"] = "post"

	c.FormParsing()
	req := c.Ctx.Request
	valid, project := validations.Projectvalidate(req)

	if !c.CheckErrors(valid, project) {
		flash := beego.NewFlash()
		_, err := o.Insert(&project)
		if err != nil {
			flash.Error("There is some problem while record insert")
		} else {
			flash.Success("Project Created Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.TplName = "projects/add-projects.tpl"
}

// EditProject edit form for project module
func (c *ProjectController) EditProject() {
	flash := beego.NewFlash()
	projectId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	project := models.Projects{Id: projectId}
	err := o.Read(&project)

	if err != nil {
		flash.Error("Project Not Found!!!")
		flash.Store(&c.Controller)
		c.Abort("404")
	}

	c.Data["Project"] = project
	c.Data["update"] = true
	c.Data["method"] = "put"
	c.TplName = "projects/add-projects.tpl"
}

// UpdateProject to update the project module
func (c *ProjectController) UpdateProject() {
	o := orm.NewOrm()
	projectId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	project := models.Projects{Id: projectId}
	o.Read(&project)
	c.FormParsing()
	req := c.Ctx.Request
	valid, updatedProject := validations.Projectvalidate(req)
	updatedProject.Id = projectId

	if !c.CheckErrors(valid, updatedProject) {
		flash := beego.NewFlash()
		_, err := o.Update(&updatedProject)
		if err != nil {
			flash.Error("There is some problem while record update")
		} else {
			flash.Success("Project Updated Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.Data["Project"] = updatedProject
	c.Data["update"] = true
	c.Data["method"] = "put"

	c.TplName = "projects/add-projects.tpl"
}

// DeleteProject to delete project
func (c *ProjectController) DeleteProject() {
	o := orm.NewOrm()
	projectId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	flash := beego.NewFlash()

	_, err := o.QueryTable(models.Projects{}).Filter("id", projectId).Update(orm.Params{
		"is_active": 0,
	})

	if err != nil {
		flash.Set("custom_error", "There is some problem while deleteing record")
	} else {
		flash.Set("custom_success", "Project Deleted Successfully!!!")
	}

	flash.Store(&c.Controller)

	c.Redirect(beego.URLFor("ProjectController.ListProjects"), 303)
}

//FormParsing checks form parsing
func (c *ProjectController) FormParsing() {
	project := models.Projects{}
	if err := c.ParseForm(&project); err != nil {
		flash := beego.NewFlash()
		flash.Error("Form Parsing Failed")
		flash.Store(&c.Controller)

		return
	}
}

//CheckErrors checks errors while validating
func (c *ProjectController) CheckErrors(valid validation.Validation, project models.Projects) bool {
	if valid.HasErrors() {
		flash := beego.NewFlash()
		flash.Error("Validation Failed!!")
		flash.Store(&c.Controller)

		c.Data["Errors"] = valid.ErrorsMap
		c.Data["Project"] = project
		return true
	}

	return false
}
