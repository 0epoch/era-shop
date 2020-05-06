package model

import "time"

type User struct {
	ID int64
	AccountID int64
	Nickname string
	Avatar string
	Bio string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
