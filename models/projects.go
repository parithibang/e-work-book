package models

import "github.com/astaxie/beego/orm"

// Projects Model Struct
type Projects struct {
	Id       int      `orm:"auto"`
	Name     string   `orm:"size(255);null" valid:"Required"`
	IsActive int      `orm:"default(1)"`
	Pods     []*Pods  `orm:"reverse(many);null;rel_through(github.com/parithibang/e-work-book/models.PodsProjects)"`
	Users    []*Users `orm:"reverse(many);null;rel_through(github.com/parithibang/e-work-book/models.UsersProjects)"`
}

// GetAllProjects list all the pods
func (project *Projects) GetAllProjects(limit, page int) ([]*Projects, int64) {
	var projects []*Projects
	o := orm.NewOrm()
	setter := o.QueryTable(Projects{}).Filter("is_active", 1).OrderBy("id")
	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&projects)

	return projects, count
}

//GetProjects list all the projects
func (project *Projects) GetProjects() []*Projects {
	var projects []*Projects
	o := orm.NewOrm()
	o.QueryTable(Projects{}).OrderBy("name").All(&projects, "id", "name")
	return projects
}
