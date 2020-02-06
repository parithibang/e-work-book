package models

import "github.com/astaxie/beego/orm"

// UsersProjects Model Struct
type UsersProjects struct {
	Id         int       `orm:"auto"`
	Projects   *Projects `orm:"rel(fk);column(project_id)"`
	Users      *Users    `orm:"rel(fk);column(user_id)"`
	Percentage float64   `orm:"column(work_percentage)"`
}

// GetUserAssignedProjects to get the list of user assigned projects
func (usersProject *UsersProjects) GetUserAssignedProjects(name string, limit int, page int) ([]*UsersProjects, int64) {
	var usersProjects []*UsersProjects
	o := orm.NewOrm()
	setter := o.QueryTable(UsersProjects{}).RelatedSel()
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.AndCond(cond.And("users__first_name__istartswith", name).Or("users__last_name__iendswith", name))

	count, _ := setter.SetCond(cond).Count()
	setter.SetCond(cond).Limit(limit, (page-1)*limit).All(&usersProjects)

	return usersProjects, count
}
