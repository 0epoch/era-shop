package adminhandle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

type RoleHandle struct {}

func (rh *RoleHandle) Create(c *gin.Context) {
	bindRole := &model.Role{}
	err := c.ShouldBind(bindRole)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}

	data := map[string]interface{}{
		"name": bindRole.Name,
		"permissions": bindRole.Permissions,
	}

	role, err := service.CreateRole(data)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	response.Success(c, role)
}

func (rh *RoleHandle) Role(c *gin.Context) {

}

func (rh *RoleHandle) Roles(c *gin.Context) {

}

func (rh *RoleHandle) Modify(c *gin.Context) {
	bindRole := &model.Role{}
	err := c.ShouldBind(bindRole)
	if err != nil {
		response.Failed(c, "修改失败！")
		return
	}

	data := map[string]interface{}{
		"id": bindRole.ID,
		"name":bindRole.Name,
		"permissions": bindRole.Permissions,
	}

	err = service.UpdateRole(data)
	if err != nil {
		response.Failed(c, "修改失败")
		return
	}
	response.Success(c, "")
}

func (rh *RoleHandle) Deleted(c *gin.Context) {

}