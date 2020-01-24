package models

import (
	"time"
)

// Teams Model Struct
type Teams struct {
	Id        int
	Name      string    `orm:"size(255)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
