package model

import "time"

const (
	AccountDeleted = -1 //账户删除
	AccountNormal = 1 //正常账户，默认
)
type Account struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Password string `json:"password"`
	LastLogin time.Time `json:"last_login"`
	LastIP string `json:"last_ip"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time
}
