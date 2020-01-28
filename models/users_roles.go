package models

// UsersRoles Model Struct
type UsersRoles struct {
	Id    int    `orm:"auto"`
	Users *Users `orm:"rel(fk);column(user_id)"`
	Roles *Roles `orm:"rel(fk);column(role_id)"`
}
