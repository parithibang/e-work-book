package models

import (
	"time"
)

// RolesPermissions Model Struct
type RolesPermissions struct {
	Id           int
	PermissionId int       `orm:"size(20)"`
	RoleId       int       `orm:"size(20)"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
}
