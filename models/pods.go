package models

import (
	"time"
)

// Pods Model Struct
type Pods struct {
	Id        int
	Name      string    `orm:"size(255)"`
	TeamId    int       `orm:"size(10)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
