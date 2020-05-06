package adminhandle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/auth"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandle struct {}

func (a *AdminHandle) Create(c *gin.Context) {
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
		response.Failed(c, "添加失败！")
		return
	}

	adminData := map[string]interface{}{
		"name": c.PostForm("name"),
		"phone": account.Phone,
		"accountID": account.ID,
		"avatar": c.PostForm("avatar"),
	}
	admin, err := service.CreateAdmin(account, adminData)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}
	//设置管理员角色
	adminRole := map[string]interface{}{
		"accountID": account.ID,
		"adminID": admin.ID,
		"roles": c.PostForm("roles"),
	}
	service.CreateAdminRole(adminRole)
	response.Success(c, admin)
}

func (a *AdminHandle) Login(c *gin.Context) {
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

func (a *AdminHandle) Admin(c *gin.Context) {

}

func (a *AdminHandle) Admins(c *gin.Context) {

}

func (a *AdminHandle) Modify(c *gin.Context) {

}

func (a *AdminHandle) Deleted(c *gin.Context) {

}

func (a *AdminHandle) Menus(c *gin.Context) {
	accountID, _ := c.Get("accountID")
	menus, err := service.FindAdminMenus(accountID.(int64))
	if err != nil {
		response.Failed(c, "查找失败！")
		return
	}
	response.Success(c, menus)
}