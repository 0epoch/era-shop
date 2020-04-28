package auth

import (
	"era-shop.xyz/era-shop/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("era_shop")

type Claims struct {
	AccountID int64 `json:"account_id"`
	jwt.StandardClaims
}

func GenerateToken(account model.Account)(string, error) {
	expire := time.Now().Add(7 * 24 *time.Hour)
	claims := &Claims{
		AccountID: account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt: time.Now().Unix(), //token发放时间
			Issuer: "era_shop", //token发放者
			Subject: "user token", //说明
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	fmt.Println(err)

	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token)(i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

//TODO 用户删除后 token处理