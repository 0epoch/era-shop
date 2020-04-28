package service

import (
	"era-shop.xyz/era-shop/model"
)

type AttrService struct {}

func (a *AttrService) FindAttrByID(id int64) (*model.Attr, error) {
	return nil, nil
}

func (a *AttrService) FindAttrs() []model.Attr {

	return nil
}

func (a *AttrService) CreateAttr(attr *model.Attr) error {
	return model.DB.Create(attr).Error
}

func (a *AttrService) UpdateAttr(attr *model.Attr) error {
	return nil
}

func (a *AttrService) DeleteAttrByID(id int64) bool {
	return true
}