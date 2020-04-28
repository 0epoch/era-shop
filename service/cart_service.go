package service

import (
	"encoding/json"
	"era-shop.xyz/era-shop/model"
)

type CartService struct {}

func (c *CartService) FindCartByID(id int64) (*model.Cart, error) {
	return nil, nil
}

func (c *CartService) FindCarts() []model.Cart {

	return nil
}

func (c *CartService) CreateCart(data map[string]interface{}) (*model.Cart, error) {
	sku := &model.Sku{}
	skuID := data["skuID"].(int64)
	model.DB.First(sku, skuID)

	skuSnapshot := model.SkuSnapshot{
		SpuName: "todo spu name",
		MainImg: sku.MainImg,
		Price: sku.Price,
		MarketPrice: sku.MarketPrice,
	}
	snapshot, err := json.Marshal(skuSnapshot)

	cart := &model.Cart{
		AccountID:   data["accountID"].(int64),
		SpuID:       data["spuID"].(int64),
		SkuID:       skuID,
		SkuSnapshot: string(snapshot),
		Quantity:    data["quantity"].(int),
		Price:       sku.Price,
	}
	err = model.DB.Create(cart).Error
	return cart, err
}

func (c *CartService) UpdateCart(cart *model.Cart) error {
	return nil
}

func (c *CartService) DeleteCartByID(id int64) bool {
	return true
}