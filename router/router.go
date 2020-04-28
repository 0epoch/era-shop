package router

import (
	"era-shop.xyz/era-shop/handle"
	"era-shop.xyz/era-shop/middleware"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
)

const apiRoot = "/api/v1"

func Load() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	u := handle.User{}

	v1 := r.Group(apiRoot)
	{
		v1.POST("/login", u.Login)
		v1.POST("/register", u.Register)
	}
	r.GET(apiRoot+"/user", middleware.Auth(), u.Show)

	//分类
	category := handle.NewCategoryHandle(&service.CategoryService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/category/:id", category.Category)
		v1.GET("/categories", category.Categories)
		v1.POST("/category", category.Create)
		v1.DELETE("/category/:id", category.Delete)
		v1.PUT("/category", category.Modify)
	}


	//品牌
	brand := handle.NewBrandHandle(&service.BrandService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("brand/:id", brand.Brand)
		v1.GET("brands", brand.Brands)
		v1.POST("brand", brand.Create)
		v1.PUT("brand", brand.Modify)
	}

	//商品SPU
	spu := handle.NewSpuHandle(&service.SpuService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/spu/:id", spu.Spu)
		v1.GET("/spu_list", spu.SpuList)
		v1.POST("/spu", spu.Create)
		v1.GET("/spu", spu.Modify)
	}
	//商品SKU
	sku := handle.NewSkuHandle(&service.SkuService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/sku/:id", sku.Sku)
		v1.GET("/sku_list", sku.SkuList)
		v1.POST("/sku", sku.Create)
		v1.GET("/sku", sku.Modify)
	}

	//SPU属性
	attr := handle.NewAttrHandle(&service.AttrService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/attr/:id", attr.Attr)
		v1.GET("/attrs", attr.Attrs)
		v1.POST("/attr", attr.Create)
		v1.GET("/attr", attr.Modify)
	}

	//SPU属性值
	attrValue := handle.NewAttrValueHandle(&service.AttrValueService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/attr_value/:id", attrValue.AttrValue)
		v1.GET("/attr_values", attrValue.AttrValues)
		v1.POST("/attr_value", attrValue.Create)
		v1.GET("/attr_value", attrValue.Modify)
	}

	//购物车
	cart := handle.NewCartHandle(&service.CartService{})
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/cart/:id", cart.Cart)
		v1.GET("/carts", cart.Carts)
		v1.POST("/cart", cart.Create)
		v1.GET("/cart", cart.Modify)
	}
	return r
}

