package middleware

import (
	"era-shop.xyz/era-shop/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : 401,
				"msg": "权限认证失败!",
			})
			c.Abort()
			return
		}

		tokenStr = tokenStr[7:]
		token, claims, err := auth.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限认证失败！",
			})
			c.Abort()
			return
		}
		accountID := claims.AccountID
		//TODO 查询用户账户是否存在
		c.Set("accountID", accountID)
		c.Next()
	}
}