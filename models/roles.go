package models

// Roles Model Struct
type Roles struct {
	Id          int            `orm:"auto"`
	Name        string         `orm:"size(255)"`
	Users       []*Users       `orm:"reverse(many);null;rel_through(github.com/parithibang/e-work-book/models.UsersRoles)"`
	Permissions []*Permissions `orm:"rel(m2m);null;rel_through(github.com/parithibang/e-work-book/models.RolesPermissions)"`
}
