package handle

import (
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

type OrderHandle struct {
	orderService *service.OrderService
}

func NewOrderHandle(s *service.OrderService) *OrderHandle {
	return &OrderHandle{orderService:s}
}

func (oh *OrderHandle) Create(c *gin.Context) {

}

func (oh *OrderHandle) Order(c *gin.Context) {

}

func (oh *OrderHandle) Orders(c *gin.Context) {

}

func (oh *OrderHandle) Modify(c *gin.Context) {

}

func (oh *OrderHandle) Delete(c *gin.Context) {

}