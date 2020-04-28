package model

import "time"

type Cart struct {
	ID int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	SpuID int64 `json:"spu_id" form:"spu_id"`
	SkuID int64 `json:"sku_id" form:"sku_id"`
	SkuSnapshot string `json:"sku_snapshot"`
	Quantity int `json:"quantity" form:"quantity"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SkuSnapshot struct {
	SpuName string `json:"spu_name"`
	Price int `json:"price"`
	MarketPrice int `json:"market_price"`
	MainImg string `json:"main_img"`
}