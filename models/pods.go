package models

import "github.com/astaxie/beego/orm"

// Pods Model Struct
type Pods struct {
	Id       int         `orm:"auto"`
	Name     string      `orm:"size(255)"`
	Users    []*Users    `orm:"reverse(many);null"`
	Projects []*Projects `orm:"rel(m2m);null;rel_through(github.com/aravindkumaremis/e-work-book/models.PodsProjects)"`
	Teams    []*Teams    `orm:"reverse(many);null"`
}

//GetPods list all the pods
func (pod *Pods) GetPods() []*Pods {
	var pods []*Pods
	o := orm.NewOrm()
	o.QueryTable(Pods{}).OrderBy("name").All(&pods, "id", "name")
	return pods
}
