package models

// RolesPermissions Model Struct
type RolesPermissions struct {
	ID          int          `orm:"pk;auto;column(id)"`
	Permissions *Permissions `orm:"rel(fk);column(permission_id)"`
	Roles       *Roles       `orm:"rel(fk);column(role_id)"`
}
