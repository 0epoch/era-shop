package service

import "era-shop.xyz/era-shop/model"

func CreateUser(account *model.Account, data map[string]interface{})(*model.User, error) {
	user := &model.User{
		Nickname:     data["nickname"].(string),
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

	user.AccountID = account.ID
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, tx.Commit().Error

}