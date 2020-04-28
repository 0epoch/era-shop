package model

import "time"

type Sku struct {
	ID int64 `json:"id"`
	SpuID int64 `json:"spu_id"`
	Attrs string `json:"attrs"`
	Stock int `json:"stock"`
	BannerImg string `json:"banner_img"`
	MainImg string `json:"main_img"`
	Price int `json:"price"`
	MarketPrice int `json:"market_price"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time
}