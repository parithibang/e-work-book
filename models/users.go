package models

import "github.com/astaxie/beego/orm"

// Users Model Struct
type Users struct {
	Id        int         `orm:"auto"`
	FirstName string      `orm:"size(255)" valid:"Required"`
	LastName  string      `orm:"size(255)"`
	UserName  string      `orm:"size(255)" valid:"Email; MaxSize(100)"`
	Password  string      `orm:"type(text)" valid:"Required;MinSize(8)"`
	IsPodLead int         `orm:"size(1);default(0)"`
	Pods      *Pods       `orm:"rel(fk);column(pod_id);null"`
	Teams     *Teams      `orm:"rel(fk);column(team_id);null"`
	Roles     []*Roles    `orm:"rel(m2m);null;rel_through(github.com/aravindkumaremis/e-work-book/models.UsersRoles)"`
	Projects  []*Projects `orm:"rel(m2m);null;rel_through(github.com/aravindkumaremis/e-work-book/models.UsersProjects)"`
}

// GetAllUsers list all the users
func (user *Users) GetAllUsers(limit, page int) ([]*Users, int64) {
	var users []*Users
	o := orm.NewOrm()
	setter := o.QueryTable(Users{}).OrderBy("id").RelatedSel("Pods")
	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&users)

	return users, count
}

// FullName list all the users
func (user *Users) FullName() string {
	return user.FirstName + " " + user.LastName
}
