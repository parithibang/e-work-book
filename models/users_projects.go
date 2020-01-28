package models

// UsersProjects Model Struct
type UsersProjects struct {
	Id       int       `orm:"auto"`
	Projects *Projects `orm:"rel(fk);column(project_id)"`
	Users    *Users    `orm:"rel(fk);column(user_id)"`
}
