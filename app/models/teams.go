package models

import "github.com/astaxie/beego/orm"

// Teams Model Struct
type Teams struct {
	ID       int    `orm:"pk;auto;column(id)"`
	Name     string `orm:"size(255)" valid:"Required"`
	IsActive int    `orm:"default(1)"`
	Pods     *Pods  `orm:"rel(fk);column(pod_id);null"`
}

// GetAllTeams list all the teams
func (team *Teams) GetAllTeams(limit, page int) ([]*Teams, int64) {
	var teams []*Teams
	o := orm.NewOrm()
	setter := o.QueryTable(Teams{}).Filter("is_active", 1).OrderBy("id").RelatedSel("Pods")
	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&teams)

	return teams, count
}
