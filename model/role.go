package model

import "time"

type Role struct {
	ID int `json:""`
	Name string `json:"name" form:"name"`
	Permissions string `json:"permissions" form:"permissions"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}