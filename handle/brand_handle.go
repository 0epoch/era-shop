package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/service"
	"era-shop.xyz/era-shop/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BrandHandle struct {
	brandService *service.BrandService
}

func NewBrandHandle(s *service.BrandService) *BrandHandle {
	return &BrandHandle{brandService:s}
}

func (bh *BrandHandle) Create(c *gin.Context) {
	brand := &model.Brand{
		Name: c.PostForm("name"),
		Logo: c.PostForm("logo"),
		Desc: c.PostForm("desc"),
	}

	err := bh.brandService.CreateBrand(brand)
	if err != nil {
		response.Failed(c, "添加失败!")
		return
	}
	response.Success(c, brand)
}

func (bh *BrandHandle) Brand(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	brand, err := bh.brandService.FindBrandByID(int64(id))
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c, brand)
}

func (bh *BrandHandle) Brands(c *gin.Context) {

}


func(bh *BrandHandle) Modify(c *gin.Context) {
	status, _ := strconv.Atoi(c.PostForm("status"))
	id, _ := strconv.Atoi(c.PostForm("id"))
	brand := &model.Brand{
		ID: int64(id),
		Name:       c.PostForm("name"),
		Logo:       c.PostForm("logo"),
		Desc:       c.PostForm("desc"),
		Status:     status,
	}
	err := bh.brandService.UpdateBrand(brand)
	if err != nil {
		response.Failed(c, "修改失败！")
		return
	}
	response.Success(c, brand)
}

func(bh *BrandHandle) Delete(c *gin.Context) {

}
