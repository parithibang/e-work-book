package models

// Permissions Model Struct
type Permissions struct {
	ID    int      `orm:"pk;auto;column(id)"`
	Name  string   `orm:"size(255)"`
	Roles []*Roles `orm:"reverse(many);null;rel_through(github.com/parithibang/e-work-book/app/models.RolesPermissions)"`
}
