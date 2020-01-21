package models

import (
	"time"
)

// Users Struct
type Users struct {
	Id      int
	Name    string    `orm:"size(100)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
