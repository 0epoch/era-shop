package handle

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/response"
	"era-shop.xyz/era-shop/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type CategoryHandle struct{
	categoryService *service.CategoryService
}

func NewCategoryHandle(s *service.CategoryService) *CategoryHandle{
	return &CategoryHandle{categoryService:s}
}

func(ch *CategoryHandle) Create(c *gin.Context) {
	pid,_ := strconv.Atoi(c.DefaultPostForm("pid", "0"))

	category := &model.Category{
		Name: c.PostForm("name"),
		Pid: int64(pid),
		Desc: c.PostForm("desc"),
		Image: c.PostForm("image"),
		Brands: ","+c.PostForm("brands")+",",
	}

	err := ch.categoryService.CreateCategory(category)
	if err != nil {
		response.Send(c, 4000, err.Error(), "")
		return
	}
	response.Success(c, category)
}

func(ch *CategoryHandle) Category(c *gin.Context) {
	id := c.Param("id")
	data := &model.Category{}
	model.DB.First(data,id)
	if data.ID <= 0 {
		response.Failed(c, "分类不存在！")
		return
	}
	response.Success(c, data)
}

//分类列表
func(ch *CategoryHandle) Categories(c *gin.Context) {
	conditions := map[string]interface{} {
		"name": c.Query("name"),
		"created_at": c.Query("created_at"),
		"updated_at": c.Query("updated_at"),
	}
	data,err := ch.categoryService.FindCategories(conditions, 1, 10)
	if err != nil {
		response.NotFond(c, "无数据")
		return
	}
	response.Success(c, data)
}

//修改分类
func(ch *CategoryHandle) Modify(c *gin.Context) {
	id := c.PostForm("id")
	category := &model.Category{}
	model.DB.First(category, id)
	if category.ID <= 0 {
		response.NotFond(c, "分类不存在！")
		return
	}
	category.Name = c.DefaultPostForm("name", category.Name)
	category.Image = c.DefaultPostForm("image", category.Image)
	category.Desc = c.DefaultPostForm("desc", category.Desc)
	category.Brands = ","+c.PostForm("brands")+","

	//TODO 状态默认值
	category.Status,_ = strconv.Atoi(c.PostForm("status"))
	err := model.DB.Save(category).Error
	if err != nil {
		response.Failed(c, "修改失败！")
		return
	}
	response.Success(c, category)
}

//删除分类
func(ch *CategoryHandle) Delete(c *gin.Context) {
	id := c.Param("id")
	spu := &model.Spu{}
	model.DB.Where("category_id = ?", spu)
	if spu.ID > 0 {
		response.Failed(c, "不可删除，已有SPU使用该分类！")
	}
	err := model.DB.Model(&model.Category{}).Where("id = ?", id).Update(map[string]interface{}{
		"deleted_at":time.Now(),
		"status":model.CategoryDeleted},
	).Error
	if err != nil {
		response.Failed(c, "删除失败！")
	}
	response.Success(c, "")
}
