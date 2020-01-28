package models

// PodsProjects Model Struct
type PodsProjects struct {
	Id       int       `orm:"auto"`
	Pods     *Pods     `orm:"rel(fk);column(pod_id)"`
	Projects *Projects `orm:"rel(fk);column(project_id)"`
}
