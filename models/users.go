package models

import (
	"time"
)

// Users Model Struct
type Users struct {
	Id        int
	Name      string    `orm:"size(255)"`
	IsPodLead int       `orm:"size(1)"`
	PodId     int       `orm:"size(20)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
