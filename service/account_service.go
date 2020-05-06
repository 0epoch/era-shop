package service

import (
	"era-shop.xyz/era-shop/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func FindAccountByID() {

}

func FindAccounts() {

}

func CreateAccount (data map[string]interface{}) (*model.Account, error) {

	pwd := data["password"].(string)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := &model.Account{
		Username: data["username"].(string),
		Phone:     data["phone"].(string),
		Email:     data["email"].(string),
		Status: model.AccountNormal,
		Password:  string(passwordHash),
		LastLogin: time.Now(),
	}
	return account, nil
}

