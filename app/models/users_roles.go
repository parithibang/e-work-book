package models

// UsersRoles Model Struct
type UsersRoles struct {
	ID    int    `orm:"pk;auto;column(id)"`
	Users *Users `orm:"rel(fk);column(user_id)"`
	Roles *Roles `orm:"rel(fk);column(role_id)"`
}
