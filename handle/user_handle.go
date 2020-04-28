package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/auth"
	"era-shop.xyz/era-shop/pkg/helper"
	"era-shop.xyz/era-shop/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)
type User struct {}

func (uh *User) Show(c *gin.Context) {
	response.Success(c, c.Param("id"))
}

//注册
func (uh *User) Register(c *gin.Context) {
	pwd := c.PostForm("password")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Failed(c, "注册失败！")
		return
	}

	account := model.Account{
		Username: helper.RandStr(8),
		Phone:     c.PostForm("phone"),
		Email:     "",
		Status: model.AccountNormal,
		Password:  string(passwordHash),
		LastLogin: time.Now(),
	}
	
	model.DB.Create(&account)
	if account.ID <= 0 {
		response.Failed(c, "注册失败！")
		return
	}

	user := model.User{
		AccountID: account.ID,
		Nickname:  account.Username,
		Avatar:    "",
		Bio:       "",
	}
	model.DB.Create(&user)
	response.Success(c, user)
}

//登录
func (uh *User) Login(c *gin.Context) {
	pwd := c.PostForm("password")
	account := model.Account{}
	model.DB.Where("phone = ?", c.PostForm("phone")).First(&account)

	if account.ID <= 0 {
		response.NotFond(c,"用户不存在！")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(pwd))
	if err != nil {
		response.Send(c, 4000,"密码错误", "")
		return
	}

	token, er := auth.GenerateToken(account)
	if er != nil {
		response.Send(c, 4001,"授权失败！", "")
		return
	}

	response.Success(c, token)
}