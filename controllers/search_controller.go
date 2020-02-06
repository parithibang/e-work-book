package controllers

import (
	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego/orm"
)

// SearchController doc
type SearchController struct {
	BaseController
}

// GetSearch to list user add form
func (c *SearchController) GetSearch() {
	o := orm.NewOrm()
	var pods []*models.Pods
	o.QueryTable(models.Pods{}).OrderBy("name").All(&pods, "id", "name")
	c.Data["pods"] = pods

	var teams []*models.Teams
	o.QueryTable(models.Teams{}).OrderBy("name").All(&teams, "id", "name")
	c.Data["teams"] = teams

	var projects []*models.Projects
	o.QueryTable(models.Projects{}).OrderBy("name").All(&projects, "id", "name")
	c.Data["projects"] = projects
	c.Data["searchmenu"] = 1
	c.TplName = "search/search.tpl"
}

// PostSearchResults create a new user
func (c *SearchController) PostSearchResults() {
	c.TplName = "search/search.tpl"
}
