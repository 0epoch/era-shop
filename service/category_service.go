package service

import (
	"era-shop.xyz/era-shop/model"
	"errors"
	"strconv"
)

type CategoryService struct {}

func(service *CategoryService) FindCategoryByID(id int64) *model.Category{

	return &model.Category{}
}

func(service *CategoryService) FindCategories(condition map[string]interface{}, page, pageSize int) (map[string]interface{},error) {
		offset := (page - 1) * pageSize
		var total int
		var categories []model.Category
		m := model.DB.Model(model.Category{})
		for k, v := range condition {
			if v == "" {
				continue
			}
			switch k {
			case "name":
				m = m.Where("name LIKE ?", "%"+v.(string)+"%")
			case "status":
				m = m.Where("status = ?", v)
			}
		}
		m.Count(&total)
		err := m.Offset(offset).Limit(pageSize).Find(&categories).Error
		if err != nil {
			return nil, err
		}
		paginate := map[string]interface{} {
			"total": total,
			"page": page,
			"page_size": pageSize,
		}

		return map[string]interface{} {
			"list": categories,
			"paginate": paginate,
		}, nil
}

func(service *CategoryService) DeleteCategoryByID(id int64) bool {

	return true
}

func(service *CategoryService) UpdateCategory() {

}

func(service *CategoryService) CreateCategory(category *model.Category) error {
	if category.Pid > 0 {
		return service.createChildCategory(category)
	}
	return model.DB.Create(category).Error
}

//添加子分类
func(service *CategoryService) createChildCategory(category *model.Category)  error {
	parentCategory := model.Category{}
	model.DB.Where("id = ?", category.Pid).First(&parentCategory)

	if parentCategory.ID <= 0 {
		return errors.New("父分类不存在")
	}

	pid := strconv.Itoa(int(parentCategory.ID))
	if parentCategory.Path == "" {
		category.Path = pid
	} else {
		category.Path = parentCategory.Path + "-" + pid
	}
	model.DB.Create(category)
	if category.ID <= 0 {
		return errors.New("添加失败")
	}
	return nil
}


