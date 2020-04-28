package service

import (
	"era-shop.xyz/era-shop/model"
	"strconv"
)

type SpuService struct {}

func (s *SpuService) FindSpuByID(id int64) *model.Spu {
	return nil
}

func (s *SpuService) FindSpuListByID() []*model.Spu {
	return nil
}

func (s *SpuService) CreateSpu(data map[string]interface{}) (*model.Spu, error) {
	floatPrice, _ := strconv.ParseFloat(data["price"].(string), 64)
	floatMarketPrice, _ := strconv.ParseFloat(data["marketPrice"].(string), 64)
	price := int(floatPrice * model.PDP)
	marketPrice := int(floatMarketPrice * model.PDP)

	spu := &model.Spu{
		Name:        data["name"].(string),
		Unit:        data["unit"].(string),
		BrandID:     data["brandID"].(int64),
		CategoryID:  data["categoryID"].(int64),
		BannerImg:   data["bannerImg"].(string),
		MainImg:     data["mainImg"].(string),
		Price:       price,
		MarketPrice: marketPrice,
		Desc:        data["desc"].(string),
	}

	err := model.DB.Create(spu).Error
	return spu, err
}

func (s *SpuService) UpdateSpu(spu *model.Spu) error {
	return nil
}

func (s *SpuService) DeleteSpuByID(id int64) error {
	return nil
}
