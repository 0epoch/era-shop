package model

import "time"

type Attr struct {
	ID int64 `json:""`
	Name string `json:""`
	Desc string `json:""`
	Status int `json:""`
	CreatedAt time.Time `json:""`
	UpdatedAt time.Time `json:""`
	DeletedAt *time.Time
}