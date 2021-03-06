package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/auth"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type User struct {}

func (uh *User) Show(c *gin.Context) {
	response.Success(c, c.Param("id"))
}

//注册
func (uh *User) Register(c *gin.Context) {
	accountBind := &model.Account{}
	err := c.ShouldBind(accountBind)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}

	data := map[string]interface{}{
		"username": accountBind.Username,
		"password": accountBind.Password,
		"phone": accountBind.Phone,
		"email": accountBind.Email,
	}
	account, err := service.CreateAccount(data)
	if err != nil {
		response.Failed(c, "注册失败！")
		return
	}

	userData := map[string]interface{}{
		"nickname": "",
		"avatar": "",
	}
	user, err := service.CreateUser(account, userData)
	if err != nil {
		response.Failed(c, "注册失败！")
		return
	}
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