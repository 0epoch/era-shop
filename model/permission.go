package model

import "time"

type Permission struct {
	ID int `json:"id"`
	Pid int `json:"pid" form:"pid"`
	Name string `json:"name" form:"name"`
	Icon string `json:"icon" form:"icon"`
	Uri string `json:"route_uri" form:"route_uri"`
	Method string `json:"method" form:"method"`
	Order int `json:"order" form:"order"`
	IsMenu int `json:"is_menu" form:"is_menu"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
