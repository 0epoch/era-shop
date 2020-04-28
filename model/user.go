package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	AccountID int64
	Nickname string
	Avatar string
	Bio string
}
