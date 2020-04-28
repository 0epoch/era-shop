package service

import (
	"era-shop.xyz/era-shop/model"
	"errors"
)

type BrandService struct {}

func (b *BrandService) FindBrandByID(id int64) (*model.Brand, error) {
	brand := &model.Brand{}
	err := model.DB.First(brand, id).Error
	if err != nil {
		return nil, errors.New("品牌不存在！")
	}
	return brand, nil
}

func (b *BrandService) FindBrands() []model.Brand {

	var brands []model.Brand
	return brands
}

func (b *BrandService) CreateBrand(brand *model.Brand) error {
	return model.DB.Create(brand).Error
}

func (b *BrandService) UpdateBrand(brand *model.Brand) error {
	err := model.DB.Save(brand).Error
	return err
}

func (b *BrandService) DeleteBrandByID(id int64) bool {
	return true
}