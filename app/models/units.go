package models

import "github.com/astaxie/beego/orm"

// Units Model Struct
type Units struct {
	Id       int     `orm:"auto"`
	Name     string  `orm:"size(255)"`
	IsActive int     `orm:"default(1)"`
	Pods     []*Pods `orm:"reverse(many);null"`
}

//GetUnits list all the units
func (unit *Units) GetUnits() []*Units {
	var units []*Units
	o := orm.NewOrm()
	o.QueryTable(Units{}).OrderBy("name").All(&units, "id", "name")
	return units
}
