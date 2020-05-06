package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

type AttrHandle struct {
	attrService *service.AttrService
}

func NewAttrHandle(s *service.AttrService) *AttrHandle {
	return &AttrHandle{attrService:s}
}

func (ah *AttrHandle) Create(c *gin.Context) {
	attr := &model.Attr{
		Name:      c.PostForm("name"),
		Desc:      c.PostForm("desc"),
	}

	err := ah.attrService.CreateAttr(attr)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	response.Success(c, attr)
}

func (ah *AttrHandle) Attr(c *gin.Context) {

}

func (ah *AttrHandle) Attrs(c *gin.Context) {

}

func (ah *AttrHandle) Modify(c *gin.Context) {

}

func (ah *AttrHandle) Delete(c *gin.Context) {

}