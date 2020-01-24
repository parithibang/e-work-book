package models

import (
	"time"
)

// Projects Model Struct
type Projects struct {
	Id        int
	Name      string    `orm:"size(255)"`
	PodId     int       `orm:"size(20)"`
	UserId    int       `orm:"size(20)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
