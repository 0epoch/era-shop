package service

import (
	"era-shop.xyz/era-shop/model"
)

type AttrValueService struct {}

func (v *AttrValueService) FindAttrValueByID(id int64) (*model.AttrValue, error) {
	return nil, nil
}

func (v *AttrValueService) FindAttrValues() []model.AttrValue {

	return nil
}

func (v *AttrValueService) CreateAttrValue(attrValue *model.AttrValue) error {
	return model.DB.Create(attrValue).Error
}

func (v *AttrValueService) UpdateAttrValue(attrValue *model.AttrValue) error {
	return nil
}

func (v *AttrValueService) DeleteAttrValueByID(id int64) bool {
	return true
}