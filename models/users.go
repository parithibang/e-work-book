package models

// Users Model Struct
type Users struct {
	Id        int         `orm:"auto"`
	Name      string      `orm:"size(255)"`
	IsPodLead int         `orm:"size(1);default(0)"`
	Pods      *Pods       `orm:"rel(fk);column(pod_id);null"`
	Roles     []*Roles    `orm:"rel(m2m);null;rel_through(github.com/aravindkumaremis/e-work-book/models.UsersRoles)"`
	Projects  []*Projects `orm:"rel(m2m);null;rel_through(github.com/aravindkumaremis/e-work-book/models.UsersProjects)"`
}
