package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/service"
	"era-shop.xyz/era-shop/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AttrValueHandle struct {
	attrValueService *service.AttrValueService
}

func NewAttrValueHandle(s *service.AttrValueService) *AttrValueHandle {
	return &AttrValueHandle{attrValueService:s}
}

func (vh *AttrValueHandle) Create(c *gin.Context) {
	attrID, _ := strconv.ParseInt(c.PostForm("attr_id"), 10, 64)
	attrValue := &model.AttrValue{
		AttrID:    attrID,
		Name:      c.PostForm("name"),
		Desc:      c.PostForm("desc"),
	}

	err := vh.attrValueService.CreateAttrValue(attrValue)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	response.Success(c, attrValue)
}

func (vh *AttrValueHandle) AttrValue(c *gin.Context) {

}

func (vh *AttrValueHandle) AttrValues(c *gin.Context) {

}

func (vh *AttrValueHandle) Modify(c *gin.Context) {

}

func (vh *AttrValueHandle) Delete(c *gin.Context) {

}