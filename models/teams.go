package models

// Teams Model Struct
type Teams struct {
	Id   int     `orm:"auto"`
	Name string  `orm:"size(255)"`
	Pods []*Pods `orm:"reverse(many);null"`
}
