package models

import (
	"github.com/astaxie/beego/orm"
)

// UsersProjects Model Struct
type UsersProjects struct {
	Id         int       `orm:"auto"`
	Projects   *Projects `orm:"rel(fk);column(project_id)"`
	IsActive   int       `orm:"TINYINT(1);default(1)"`
	Users      *Users    `orm:"rel(fk);column(user_id)"`
	Percentage float64   `orm:"column(work_percentage)" valid:"Required"`
}

// GetUserAssignedProjects to get the list of user assigned projects
func (usersProject *UsersProjects) GetUserAssignedProjects(name string) []*UsersProjects {
	var usersProjects []*UsersProjects
	o := orm.NewOrm()
	setter := o.QueryTable(UsersProjects{}).RelatedSel()
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.AndCond(cond.And("users__first_name__istartswith", name).Or("users__last_name__iendswith", name))

	setter.SetCond(cond).All(&usersProjects)

	return usersProjects
}

// GetTotalWorkPercentageOfUser to get the list of user assigned projects
func (usersProject *UsersProjects) GetTotalWorkPercentageOfUser(userId int) orm.ParamsList {
	o := orm.NewOrm()
	var list orm.ParamsList
	o.Raw("SELECT SUM(work_percentage) AS percentage FROM wb_users_projects where user_id =?", userId).ValuesFlat(&list)

	return list
}
