package model

import "time"

type Brand struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Desc string `json:"desc"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}