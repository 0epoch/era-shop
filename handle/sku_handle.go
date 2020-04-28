package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/service"
	"era-shop.xyz/era-shop/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SkuHandle struct {
	skuService *service.SkuService
}

func NewSkuHandle(s *service.SkuService) *SkuHandle {
	return &SkuHandle{skuService:s}
}

func (sh *SkuHandle) Create(c *gin.Context) {
	spuID, _ := strconv.ParseInt(c.PostForm("spu_id"), 10, 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))

	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	marketPrice, _ := strconv.ParseFloat(c.PostForm("market_price"), 64)

	sku := &model.Sku{
		SpuID:       spuID,
		Attrs:       c.PostForm("attrs"),
		Stock:       stock,
		BannerImg:   c.PostForm("banner_img"),
		MainImg:     c.PostForm("main_img"),
		Price:       int(price*100.0),
		MarketPrice: int(marketPrice*100.0),
	}

	err := sh.skuService.CreateSku(sku)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	response.Success(c, sku)
}

func (sh *SkuHandle) Sku(c *gin.Context) {

}

func (sh *SkuHandle) SkuList(c *gin.Context) {

}

func (sh *SkuHandle) Modify(c *gin.Context) {

}

func (sh *SkuHandle) Delete(c *gin.Context) {

}