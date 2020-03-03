package models

import "github.com/astaxie/beego/orm"

// Pods Model Struct
type Pods struct {
	Id       int         `orm:"auto"`
	Name     string      `orm:"size(255)" valid:"Required"`
	Users    []*Users    `orm:"reverse(many);null"`
	IsActive int         `orm:"default(1)"`
	Projects []*Projects `orm:"rel(m2m);null;rel_through(github.com/parithibang/e-work-book/models.PodsProjects)"`
	Teams    []*Teams    `orm:"reverse(many);null"`
	Units    *Units      `orm:"rel(fk);column(unit_id);null"`
}

//GetPods list all the pods
func (pod *Pods) GetPods() []*Pods {
	var pods []*Pods
	o := orm.NewOrm()
	o.QueryTable(Pods{}).OrderBy("name").All(&pods, "id", "name")
	return pods
}

// GetAllPods list all the pods
func (pod *Pods) GetAllPods(limit, page int) ([]*Pods, int64) {
	var teams []*Pods
	o := orm.NewOrm()
	setter := o.QueryTable(Pods{}).Filter("is_active", 1).OrderBy("id").RelatedSel("Units")
	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&teams)

	return teams, count
}
