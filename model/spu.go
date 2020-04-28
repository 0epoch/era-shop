package model

import "time"

const (
	PDP = 100.0 //price decimal places 价格小数位数
)

type Spu struct {
	ID int64 `json:"id"`
	Name string `json:"name" form:"name"`
	Unit string `json:"unit" form:"unit"`
	BrandID int64 `json:"brand_id" form:"brand_id"`
	CategoryID int64 `json:"category_id" form:"category_id"`
	BannerImg string `json:"banner_img" form:"banner_img"`
	MainImg string `json:"main_img" form:"main_img"`
	Price int `json:"price"`
	MarketPrice int `json:"market_price"`
	Desc string `json:"desc" form:"desc"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
