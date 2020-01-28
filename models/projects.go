package models

// Projects Model Struct
type Projects struct {
	Id    int      `orm:"auto"`
	Name  string   `orm:"size(255);null"`
	Pods  []*Pods  `orm:"reverse(many);null;rel_through(github.com/aravindkumaremis/e-work-book/models.PodsProjects)"`
	Users []*Users `orm:"reverse(many);null;rel_through(github.com/aravindkumaremis/e-work-book/models.UsersProjects)"`
}
