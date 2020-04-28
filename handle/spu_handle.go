package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

type SpuHandle struct {
	spuService *service.SpuService
}

func NewSpuHandle(s *service.SpuService) *SpuHandle {
	return &SpuHandle{spuService:s}
}

func (sh *SpuHandle) Create(c *gin.Context) {
	spuBind := &model.Spu{}
	err := c.ShouldBind(spuBind)
	if err != nil {
		response.Failed(c, "添加失败！")
	}

	data := map[string]interface{} {
		"name": spuBind.Name,
		"bannerImg": spuBind.BannerImg,
		"mainImg": spuBind.MainImg,
		"brandID": spuBind.BrandID,
		"categoryID": spuBind.CategoryID,
		"price": c.PostForm("price"),
		"marketPrice": c.PostForm("market_price"),
		"desc": spuBind.Desc,
		"unit": spuBind.Unit,
	}

	spu, err := sh.spuService.CreateSpu(data)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	response.Success(c, spu)
}

func (sh *SpuHandle) Spu(c *gin.Context) {

}

func (sh *SpuHandle) SpuList(c *gin.Context) {

}

func (sh *SpuHandle) Modify(c *gin.Context) {

}

func (sh *SpuHandle) Delete(c *gin.Context) {

}