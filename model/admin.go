package model

import "time"

type Admin struct {
	ID int `json:"id"`
	AccountID int64 `json:"account_id"`
	Name string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Avatar string `json:"avatar" form:"avatar"`
	Role AdminRole `json:"admin_role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}