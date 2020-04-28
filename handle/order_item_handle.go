package handle

import (
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

type OrderItemHandle struct {
	orderItemService *service.OrderItemService
}

func NewOrderItemHandle(s *service.OrderItemService) *OrderItemHandle {
	return &OrderItemHandle{orderItemService:s}
}

func (oih *OrderItemHandle) Create(c *gin.Context) {

}

func (oih *OrderItemHandle) OrderItem(c *gin.Context) {

}

func (oih *OrderItemHandle) OrderItems(c *gin.Context) {

}

func (oih *OrderItemHandle) Modify(c *gin.Context) {

}

func (oih *OrderItemHandle) Delete(c *gin.Context) {

}