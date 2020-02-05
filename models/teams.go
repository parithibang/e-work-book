package models

// Teams Model Struct
type Teams struct {
	Id    int      `orm:"auto"`
	Name  string   `orm:"size(255)"`
	Pods  *Pods    `orm:"rel(fk);column(pod_id);null"`
	Users []*Users `orm:"reverse(many);null"`
}
