package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
	"github.com/parithibang/e-work-book/app/validations"
)

// PodController Controller actions for pods
type PodController struct {
	BaseController
}

// Prepare teams controller
func (c *PodController) Prepare() {
	c.BaseController.Prepare()
	c.Data["Title"] = "Pods"
	c.Data["podsMenu"] = 1
}

// ListPods to list all the pods
func (c *PodController) ListPods() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "pods/list-pods.tpl"

	currentPage := 1
	if page, err := strconv.Atoi(c.Input().Get("p")); err == nil {
		currentPage = page
	}

	pods := models.Pods{}
	podList, count := pods.GetAllPods(pageLimit, currentPage)

	pagination.SetPaginator(c.Ctx, pageLimit, count)
	pageStart := currentPage*pageLimit - (pageLimit - 1)

	c.Data["podList"] = podList
	c.Data["deleteMethod"] = "delete"
	c.Data["pageStart"] = pageStart
	c.Data["count"] = count
}

// AddPod to list team add form
func (c *PodController) AddPod() {
	units := new(models.Units).GetUnits()
	c.Data["units"] = units
	c.Data["create"] = true
	c.Data["method"] = "post"
	c.TplName = "pods/add-pods.tpl"
}

// CreatePod to list team add form
func (c *PodController) CreatePod() {
	o := orm.NewOrm()
	units := new(models.Units).GetUnits()
	c.Data["units"] = units
	c.Data["create"] = true
	c.Data["method"] = "post"

	c.FormParsing()
	req := c.Ctx.Request
	valid, pod := validations.Podvalidate(req)

	if !c.CheckErrors(valid, pod) {
		flash := beego.NewFlash()
		_, err := o.Insert(&pod)
		if err != nil {
			flash.Error("There is some problem while record insert")
		} else {
			flash.Success("Pod Created Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.TplName = "pods/add-pods.tpl"
}

// EditPod edit form for team module
func (c *PodController) EditPod() {
	flash := beego.NewFlash()
	podID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	pod := models.Pods{ID: podID}
	err := o.Read(&pod)

	if err != nil {
		flash.Error("Pod Not Found!!!")
		flash.Store(&c.Controller)
		c.Abort("404")
	}

	units := new(models.Units).GetUnits()
	c.Data["Pod"] = pod
	c.Data["units"] = units
	c.Data["update"] = true
	c.Data["method"] = "put"
	c.TplName = "pods/add-pods.tpl"
}

// UpdatePod to update the pod module
func (c *PodController) UpdatePod() {
	o := orm.NewOrm()
	podID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	pod := models.Pods{ID: podID}
	o.Read(&pod)
	units := new(models.Units).GetUnits()
	c.FormParsing()
	req := c.Ctx.Request
	valid, updatedPod := validations.Podvalidate(req)
	updatedPod.ID = podID

	if !c.CheckErrors(valid, updatedPod) {
		flash := beego.NewFlash()
		_, err := o.Update(&updatedPod)
		if err != nil {
			flash.Error("There is some problem while record update")
		} else {
			flash.Success("Pod Updated Successfully!!!")
		}
		flash.Store(&c.Controller)
	}

	c.Data["Pod"] = updatedPod
	c.Data["units"] = units
	c.Data["update"] = true
	c.Data["method"] = "put"

	c.TplName = "pods/add-pods.tpl"
}

// DeletePod to delete pod
func (c *PodController) DeletePod() {
	o := orm.NewOrm()
	podID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	flash := beego.NewFlash()

	_, err := o.QueryTable(models.Pods{}).Filter("id", podID).Update(orm.Params{
		"is_active": 0,
	})

	if err != nil {
		flash.Set("custom_error", "There is some problem while deleteing record")
	} else {
		flash.Set("custom_success", "Pod Deleted Successfully!!!")
	}

	flash.Store(&c.Controller)

	c.Redirect(beego.URLFor("PodController.ListPods"), 303)
}

//FormParsing checks form parsing
func (c *PodController) FormParsing() {
	pod := models.Pods{}
	if err := c.ParseForm(&pod); err != nil {
		flash := beego.NewFlash()
		flash.Error("Form Parsing Failed")
		flash.Store(&c.Controller)

		return
	}
}

//CheckErrors checks errors while validating
func (c *PodController) CheckErrors(valid validation.Validation, pod models.Pods) bool {
	if valid.HasErrors() {
		flash := beego.NewFlash()
		flash.Error("Validation Failed!!")
		flash.Store(&c.Controller)

		c.Data["Errors"] = valid.ErrorsMap
		c.Data["Pod"] = pod
		return true
	}

	return false
}
