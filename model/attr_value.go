package model

import "time"

type AttrValue struct {
	ID int64 `json:"id"`
	AttrID int64 `json:"attr_id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time
}