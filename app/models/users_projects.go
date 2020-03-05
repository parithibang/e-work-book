package models

import (
	"github.com/astaxie/beego/orm"
)

// UsersProjects Model Struct
type UsersProjects struct {
	ID         int       `orm:"pk;auto;column(id)"`
	Projects   *Projects `orm:"rel(fk);column(project_id)"`
	IsActive   int       `orm:"TINYINT(1);default(1)"`
	Users      *Users    `orm:"rel(fk);column(user_id)"`
	Percentage float64   `orm:"column(work_percentage)" valid:"Required"`
}

// GetUserAssignedProjects to get the list of user assigned projects
func (usersProject *UsersProjects) GetUserAssignedProjects(name string, limit, page int) ([]*UsersProjects, int64) {
	var usersProjects []*UsersProjects
	o := orm.NewOrm()
	setter := o.QueryTable(UsersProjects{}).RelatedSel()
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.AndCond(cond.And("users__first_name__istartswith", name).Or("users__last_name__iendswith", name)).And("is_active", 1)

	count, _ := setter.SetCond(cond).Count()
	// setter.SetCond(cond).All(&usersProjects)

	setter.SetCond(cond).Limit(limit, (page-1)*limit).All(&usersProjects)

	return usersProjects, count
}

// GetTotalWorkPercentageOfUser to get the list of user assigned projects
func (usersProject *UsersProjects) GetTotalWorkPercentageOfUser(userID int) orm.ParamsList {
	o := orm.NewOrm()
	var list orm.ParamsList
	o.Raw("SELECT SUM(work_percentage) AS percentage FROM wb_users_projects where user_id =? and is_active=?", userID, 1).ValuesFlat(&list)

	return list
}
