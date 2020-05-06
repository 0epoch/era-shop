package service

import (
	"era-shop.xyz/era-shop/model"
)


func FindAdminRoleByAccountID (accountID int64) (*model.AdminRole, error) {
	adminRole := &model.AdminRole{}
	err := model.DB.Where("account_id = ?", accountID).First(adminRole).Error
	return adminRole,err
}

func FindAdminRoles(conditions map[string]interface{}) []model.AdminRole {
	return nil
}

func CreateAdminRole(data map[string]interface{}) {
	adminRole := &model.AdminRole{
		AccountID: data["accountID"].(int64),
		AdminID:   data["adminID"].(int),
		Roles:     data["roles"].(string),
	}
	model.DB.Create(adminRole)
}

func UpdateAdminRole() {

}

func DeleteAdminRole() {

}