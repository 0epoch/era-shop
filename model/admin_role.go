package model

import "time"

type AdminRole struct {
	ID int `json:"id"`
	AccountID int64 `json:"account_id"`
	AdminID int `json:"admin_id"`
	Roles string `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}