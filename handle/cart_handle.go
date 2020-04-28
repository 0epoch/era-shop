package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

type CartHandle struct {
	cartService *service.CartService
}

func NewCartHandle(s *service.CartService) *CartHandle {
	return &CartHandle{cartService:s}
}

func (ch *CartHandle) Create(c *gin.Context) {
	accountID, _ := c.Get("accountID")
	accountID = accountID.(int64)

	cartBind := &model.Cart{}
	err := c.ShouldBind(cartBind)
	if err != nil {
		response.Failed(c, "加入购物车失败！")
		return
	}
	data := map[string]interface{}{
		"accountID": accountID,
		"spuID": cartBind.SkuID,
		"skuID": cartBind.SkuID,
		"quantity": cartBind.Quantity,
	}

	cart, err := ch.cartService.CreateCart(data)
	if err != nil {
		response.Failed(c, "加入购物车失败！")
		return
	}
	response.Success(c, cart)
}

func (ch *CartHandle) Cart(c *gin.Context) {

}

func (ch *CartHandle) Carts(c *gin.Context) {

}

func (ch *CartHandle) Modify(c *gin.Context) {

}

func (ch *CartHandle) Delete(c *gin.Context) {

}