package models

// PodsProjects Model Struct
type PodsProjects struct {
	ID       int       `orm:"pk;auto;column(id)"`
	Pods     *Pods     `orm:"rel(fk);column(pod_id)"`
	Projects *Projects `orm:"rel(fk);column(project_id)"`
}
