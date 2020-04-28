package model

import (
	"time"
)

const (
	CategoryDeleted = -1
	CategoryNormal = 1
)

type Category struct {
	ID int64 `json:"id"`
	Brands string `json:"brands"`
	Name string `json:"name"`
	Pid int64 `json:"pid"`
	Image string `json:"image"`
	Desc string `json:"desc"`
	Path string `json:"path"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
