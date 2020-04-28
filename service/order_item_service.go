package service

import (
	"era-shop.xyz/era-shop/model"
)

type OrderItemService struct {}

func (oi *OrderItemService) FindOrderItemByID(id int64) (*model.OrderItem, error) {
	return nil, nil
}

func (oi *OrderItemService) FindOrderItems() []model.OrderItem {

	return nil
}

func (oi *OrderItemService) CreateOrderItem(orderItem *model.OrderItem) error {
	return nil
}

func (oi *OrderItemService) UpdateOrderItem(orderItem *model.OrderItem) error {
	return nil
}

func (oi *OrderItemService) DeleteOrderItemByID(id int64) bool {
	return true
}