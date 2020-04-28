package model

type OrderItem struct {
	ID int64 `json:"id"`
	OrderNo int `json:"order_no"`
	SpuID int64 `json:"spu_id"`
	SkuID int64 `json:"sku_id"`
	Quantity int `json:"quantity"`
	RevisesTotal int `json:"revises_total"`
	Total int `json:"total"`
	Price int `json:"price"`
	rest string `json:"rest"`
}