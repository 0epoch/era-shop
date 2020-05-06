package router

import (
	"era-shop.xyz/era-shop/handle/adminhandle"
	"era-shop.xyz/era-shop/middleware"
	"github.com/gin-gonic/gin"
)

const adminApiRoot = "/api/v1/admin"

func AdminRouter(r *gin.Engine) {
	//管理员
	admin := adminhandle.AdminHandle{}
	api := r.Group(adminApiRoot, middleware.Auth(), middleware.Permission())
	{
		api.GET("/admin", admin.Admin)
		api.GET("/admins", admin.Admins)
		api.GET("/menus", admin.Menus)
		api.POST("/admin", admin.Create)
		api.POST("/login", admin.Login)
		api.PUT("/admin", admin.Modify)
		api.DELETE("/admin/:id", admin.Deleted)
	}

	//权限控制
	authRole := adminhandle.RoleHandle{}
	permission := adminhandle.PermissionHandle{}
	api = r.Group(adminApiRoot, middleware.Auth(), middleware.Permission())
	{
		//角色
		api.GET("/role", authRole.Role)
		api.GET("/roles", authRole.Roles)
		api.POST("/role", authRole.Create)
		api.PUT("/role", authRole.Modify)
		api.DELETE("/role", authRole.Deleted)

		//权限
		api.GET("/permission", permission.Permission)
		api.GET("/permissions", permission.Permissions)
		api.POST("/permission", permission.Create)
		api.PUT("/permission", permission.Modify)
		api.DELETE("/permission", permission.Delete)
	}
}