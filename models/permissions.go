package models

// Permissions Model Struct
type Permissions struct {
	Id    int      `orm:"auto"`
	Name  string   `orm:"size(255)"`
	Roles []*Roles `orm:"reverse(many);null;rel_through(github.com/aravindkumaremis/e-work-book/models.RolesPermissions)"`
}
