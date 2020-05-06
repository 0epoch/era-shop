package service

import (
	"era-shop.xyz/era-shop/model"
)


func FindAdminByID (id int) (*model.Admin, error) {
	admin := &model.Admin{
		ID:        id,
	}
	err := model.DB.First(admin).Error
	return admin, err
}

func FindAdminByAccountID(accountID int64) (*model.Admin, error) {
	admin := &model.Admin{}
	err := model.DB.Where("account_id = ?", accountID).Preload("Role").First(admin).Error
	return admin, err
}

func FindAdmins(conditions map[string]interface{}) []model.Admin {
	return nil
}

func CreateAdmin(account *model.Account, data map[string]interface{}) (*model.Admin, error) {
	admin := &model.Admin{
		Name:      data["name"].(string),
		Phone:     data["phone"].(string),
		Avatar:    data["avatar"].(string),
	}

	tx := model.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	admin.AccountID = account.ID
	if err := tx.Create(admin).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return admin, tx.Commit().Error
}

func UpdateAdmin() {

}

func DeleteAdmin() {

}

//管理员是否拥有某个权限
func HaveThisPermission(api, method string, admin *model.Admin) bool {
	//当前路由是否设置权限
	conditions := map[string]interface{}{
		"uri": api,
		"method": method,
	}
	_, err := FindPermission(conditions)
	//当前路由未设置权限 -- 不限制
	if err != nil {
		return true
	}

	//管理员未设置角色
	if admin.Role.ID <= 0 {
		return false
	}

	//获取角色下的所有权限
	roles, err := FindRolesByID(admin.Role.Roles)
	if err != nil {
		return false
	}
	var permissionsID string
	for _, v := range roles {
		permissionsID += "," + v.Permissions
	}
	//拥有的全部权限
	permissions, err := FindPermissionsByID(permissionsID)
	if err != nil {
		return false
	}

	for _, v := range permissions {
		if v.Uri == api && v.Method == method {
			return true
		}
	}
	return false
}