package models

// RolesPermissions Model Struct
type RolesPermissions struct {
	Id          int          `orm:"auto"`
	Permissions *Permissions `orm:"rel(fk);column(permission_id)"`
	Roles       *Roles       `orm:"rel(fk);column(role_id)"`
}
