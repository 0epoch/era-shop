package middleware

import (
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		api := c.FullPath()
		method := c.Request.Method

		accountID, _ := c.Get("accountID")
		admin, err := service.FindAdminByAccountID(accountID.(int64))
		if err != nil {
			response.Failed(c, "非法请求！")
			c.Abort()
			return
		}

		if ! service.HaveThisPermission(api, method, admin){
			response.Failed(c, "无权访问！")
			c.Abort()
			return
		}
		c.Next()
	}
}