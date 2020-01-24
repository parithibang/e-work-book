package models

import (
	"time"
)

// Permissions Model Struct
type Permissions struct {
	Id        int
	Name      string    `orm:"size(255)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
