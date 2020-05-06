package adminhandle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PermissionHandle struct {}

func (ph *PermissionHandle) Create(c *gin.Context) {
	bindPermission := &model.Permission{}
	err := c.ShouldBind(bindPermission)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}

	data := map[string]interface{}{
		"pid": bindPermission.Pid,
		"name": bindPermission.Name,
		"icon": bindPermission.Icon,
		"uri": bindPermission.Uri,
		"method": bindPermission.Method,
		"isMenu": bindPermission.IsMenu,
		"order": bindPermission.Order,
	}
	permission, err := service.CreatePermission(data)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	response.Success(c, permission)
}

func (ph *PermissionHandle) Permission(c *gin.Context) {

}

func (ph *PermissionHandle) Permissions(c *gin.Context) {
	nodes, err := service.FindPermissions(map[string]interface{}{})
	if err != nil {
		response.Failed(c, "无数据")
		return
	}
	response.Success(c, nodes)
}

func (ph *PermissionHandle) Modify(c *gin.Context) {
	fmt.Println("permission modify")
}

func (ph *PermissionHandle) Delete(c *gin.Context) {

}