package controllers

import (
	"strconv"

	"github.com/parithibang/e-work-book/app/models"
	"github.com/parithibang/e-work-book/app/validations"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego/validation"
)

// TeamController doc
type TeamController struct {
	BaseController
}

// Prepare teams controller
func (c *TeamController) Prepare() {
	c.BaseController.Prepare()
	c.Data["Title"] = "Teams"
	c.Data["teamsMenu"] = 1
}

// ListTeams to list all the teams
func (c *TeamController) ListTeams() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "teams/list-teams.tpl"

	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	teams := models.Teams{}
	teamList, count := teams.GetAllTeams(pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)
	pageStart := currentPage*pageLimit - (pageLimit - 1)

	c.Data["teamList"] = teamList
	c.Data["deleteMethod"] = "delete"
	c.Data["pageStart"] = pageStart
	c.Data["count"] = count
}

// AddTeam to list team add form
func (c *TeamController) AddTeam() {
	pods := new(models.Pods).GetPods()
	c.Data["pods"] = pods
	c.Data["create"] = true
	c.Data["method"] = "post"
	c.TplName = "teams/add-teams.tpl"
}

// CreateTeam to list team add form
func (c *TeamController) CreateTeam() {
	o := orm.NewOrm()
	pods := new(models.Pods).GetPods()
	c.Data["pods"] = pods
	c.Data["create"] = true
	c.Data["method"] = "post"

	c.FormParsing()
	req := c.Ctx.Request
	valid, team := validations.Teamvalidate(req)

	if !c.CheckErrors(valid, team) {
		flash := beego.NewFlash()
		_, err := o.Insert(&team)
		if err != nil {
			flash.Error("There is some problem while record insert")
		} else {
			flash.Success("Team Created Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.TplName = "teams/add-teams.tpl"
}

// EditTeam edit form for team module
func (c *TeamController) EditTeam() {
	flash := beego.NewFlash()
	teamId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	team := models.Teams{Id: teamId}
	err := o.Read(&team)

	if err != nil {
		flash.Error("Team Not Found!!!")
		flash.Store(&c.Controller)
		c.Abort("404")
	}

	pods := new(models.Pods).GetPods()
	c.Data["Team"] = team
	c.Data["pods"] = pods
	c.Data["update"] = true
	c.Data["method"] = "put"
	c.TplName = "teams/add-teams.tpl"
}

// UpdateTeam to update the team module
func (c *TeamController) UpdateTeam() {
	o := orm.NewOrm()
	teamId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	team := models.Teams{Id: teamId}
	o.Read(&team)
	pods := new(models.Pods).GetPods()
	c.FormParsing()
	req := c.Ctx.Request
	valid, updatedTeam := validations.Teamvalidate(req)
	updatedTeam.Id = teamId

	if !c.CheckErrors(valid, updatedTeam) {
		flash := beego.NewFlash()
		_, err := o.Update(&updatedTeam)
		if err != nil {
			flash.Error("There is some problem while record update")
		} else {
			flash.Success("Team Updated Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.Data["Team"] = updatedTeam
	c.Data["pods"] = pods
	c.Data["update"] = true
	c.Data["method"] = "put"

	c.TplName = "teams/add-teams.tpl"
}

// DeleteTeam to delete team
func (c *TeamController) DeleteTeam() {
	o := orm.NewOrm()
	teamId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	flash := beego.NewFlash()

	_, err := o.QueryTable(models.Teams{}).Filter("id", teamId).Update(orm.Params{
		"is_active": 0,
	})

	if err != nil {
		flash.Set("custom_error", "There is some problem while deleteing record")
	} else {
		flash.Set("custom_success", "Team Deleted Successfully!!!")
	}

	flash.Store(&c.Controller)

	c.Redirect(beego.URLFor("TeamController.ListTeams"), 303)
}

//FormParsing checks form parsing
func (c *TeamController) FormParsing() {
	team := models.Teams{}
	if err := c.ParseForm(&team); err != nil {
		flash := beego.NewFlash()
		flash.Error("Form Parsing Failed")
		flash.Store(&c.Controller)

		return
	}
}

//CheckErrors checks errors while validating
func (c *TeamController) CheckErrors(valid validation.Validation, team models.Teams) bool {
	if valid.HasErrors() {
		flash := beego.NewFlash()
		flash.Error("Validation Failed!!")
		flash.Store(&c.Controller)

		c.Data["Errors"] = valid.ErrorsMap
		c.Data["Team"] = team
		return true
	}

	return false
}
