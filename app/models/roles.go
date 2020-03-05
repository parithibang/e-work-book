package models

// Roles Model Struct
type Roles struct {
	ID          int            `orm:"pk;auto;column(id)"`
	Name        string         `orm:"size(255)"`
	Users       []*Users       `orm:"reverse(many);null;rel_through(github.com/parithibang/e-work-book/app/models.UsersRoles)"`
	Permissions []*Permissions `orm:"rel(m2m);null;rel_through(github.com/parithibang/e-work-book/app/models.RolesPermissions)"`
}
