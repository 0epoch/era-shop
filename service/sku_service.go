package service

import "era-shop.xyz/era-shop/model"

type SkuService struct {}

func (s *SkuService) FindSkuByID(id int64) *model.Sku {
	return nil
}

func (s *SkuService) FindSkuListByID() []*model.Sku {

	return nil
}

func (s *SkuService) CreateSku(spu *model.Sku) error {
	return model.DB.Create(spu).Error
}

func (s *SkuService) UpdateSku(spu *model.Spu) error {
	return nil
}

func (s *SkuService) DeleteSkuByID(id int64) error {
	return nil
}
