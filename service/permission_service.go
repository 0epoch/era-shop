package service

import (
	"era-shop.xyz/era-shop/model"
	"strings"
)

type permissionNode struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Sub []*permissionNode `json:"sub"`
}

func FindPermissionByID (id int) *model.Permission {
	return nil
}

func FindPermissionsByID(permissionsID string) ([]*model.Permission, error) {
	permissionsID = strings.Trim(permissionsID, ",")
	id := strings.Split(permissionsID, ",")
	var permissions []*model.Permission
	err := model.DB.Where("id in (?)", id).Find(&permissions).Error
	return permissions, err
}

func FindPermission(conditions map[string]interface{}) (*model.Permission, error) {
	permission := &model.Permission{}
	for k, v := range conditions {
		switch k {
		case "id":
			permission.ID = v.(int)
		case "uri":
			permission.Uri = v.(string)
		case "method":
			permission.Method = v.(string)
		}
	}
	err := model.DB.Where(permission).First(&permission).Error
	return permission, err
}


func FindPermissions(conditions map[string]interface{}) ([]*permissionNode, error){
	var permissions []*model.Permission

	m := model.DB.Model(model.Permission{})
	for k, v := range conditions {
		switch k {
			case "id":
				m = m.Where("id = ?", v)
			case "permissionsID":
				m = m.Where("id in (?)", v)
			case "uri":
				m = m.Where("uri = ?", v)
			case "method":
				m = m.Where("method = ?", v)
			case "isMenu":
				m = m.Where("is_menu = ?", v)
		}
	}
	err := model.DB.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	nodes := permissionNodeMap(permissions, 0)
	return nodes, nil
}

func CreatePermission(data map[string]interface{}) (*model.Permission, error){
	permission := &model.Permission{
		Pid:       data["pid"].(int),
		Name:      data["name"].(string),
		Icon:      data["icon"].(string),
		Uri:       data["uri"].(string),
		Method:    data["method"].(string),
		Order:     data["order"].(int),
		IsMenu:    data["isMenu"].(int),
	}
	err := model.DB.Create(permission).Error
	return permission,err
}

func UpdatePermission() {

}

func DeletePermission() {

}

func IsHavePermission(uri, method string, permissions []*model.Permission) bool {
	for _, v := range permissions {
		if v.Uri == uri && v.Method == method {
			return true
		}
	}
	return false
}

func permissionNodeMap (permissions []*model.Permission, pid int) []*permissionNode {
	var nodeMap []*permissionNode
	for _, v := range permissions {
		if  v.Pid == pid {
			node := &permissionNode{
				ID:   v.ID,
				Name: v.Name,
			}
			sub := permissionNodeMap(permissions, v.ID)
			node.Sub = sub
			nodeMap = append(nodeMap, node)
		}
	}
	return nodeMap
}

