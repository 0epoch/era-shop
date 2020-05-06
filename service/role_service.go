package service

import (
	"era-shop.xyz/era-shop/model"
	"strings"
)


func FindRoleByID (id int) *model.Role {
	return nil
}

func FindRolesByID(rolesID string) ([]*model.Role, error) {
	id := strings.Split(rolesID, ",")
	var roles []*model.Role
	err := model.DB.Where("id in (?)", id).Find(&roles).Error
	return roles, err
}

func FindRoles(conditions map[string]interface{}) []model.Role {
	return nil
}

func FindRolesPermissions(rolesID string, conditions map[string]interface{}) ([]*permissionNode, error){
	roles, err := FindRolesByID(rolesID)
	if err != nil {
		return nil, err
	}
	var permissionsID string
	for _, v := range roles {
		permissionsID += "," + v.Permissions
	}
	conditions["permissionsID"] = permissionsID
	nodes, err := FindPermissions(conditions)
	return nodes, err
}

func CreateRole(data map[string]interface{}) (*model.Role, error){
	role := &model.Role{
		Name:        data["name"].(string),
		Permissions: data["permissions"].(string),
	}

	err := model.DB.Create(role).Error
	return role, err
}

func UpdateRole(data map[string]interface{}) error {
	role := model.Role{
		ID: data["id"].(int),
	}
	err := model.DB.First(&role).Error
	if err != nil {
		return err
	}

	err = model.DB.Model(&role).Updates(data).Error
	return err
}

func DeleteRole() {

}
