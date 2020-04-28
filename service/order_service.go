package service

import (
	"era-shop.xyz/era-shop/model"
)

type OrderService struct {}

func (o *OrderService) FindOrderByID(id int64) (*model.Order, error) {
	return nil, nil
}

func (o *OrderService) FindOrders() []model.Order {

	return nil
}

func (o *OrderService) CreateOrder(order *model.Order) (*model.Order, error) {
	return nil, nil
}

func (o *OrderService) UpdateOrder(order *model.Order) error {
	return nil
}

func (o *OrderService) DeleteOrderByID(id int64) bool {
	return true
}