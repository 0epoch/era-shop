package model

import "time"

type Order struct {
	ID int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	ProductsNum int `json:"products_num"`
	Total int `json:"total"`
	PayTotal int `json:"pay_total"`
	RevisesTotal int `json:"revises_total"`
	PayStatus int `json:"pay_status"`
	ShipStatus int `json:"ship_status"`
	UserIP string `json:"user_ip"`
	PaidAt time.Time `json:"paid_at"`
	ShippedAt time.Time `json:"shipped_at"`
	ConfirmedAt time.Time `json:"confirmed_at"`
	ReviewedAt time.Time `json:"reviewed_at"`
	FinishedAt time.Time `json:"finished_at"`
}