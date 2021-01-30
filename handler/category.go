package handler

import (
	"context"

	"github.com/itgamerskyHuang/category/domain/model"
	"github.com/itgamerskyHuang/category/domain/service"
	pb "github.com/itgamerskyHuang/category/proto/category"
	"github.com/jinzhu/gorm"
)

type Category struct {
	CategoryDataService service.ICategoryService
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{
		CategoryDataService: service.NewCategoryService(db),
	}
}

func (c *Category) CreateCategory(ctx context.Context, r *pb.CreateRequest, w *pb.CreateResponse) error {
	category := &model.Category{
		CategoryName:        r.CategoryName,
		CategoryLevel:       r.CategoryLevel,
		CategoryParent:      r.CategoryParent,
		CategoryImage:       r.CategoryImage,
		CategoryDescription: r.CategoryDescription,
	}
	err := c.CategoryDataService.CreateCategory(category)
	if err != nil {
		w.Message = "创建失败"
	}
	w.Message = "创建成功"
	w.CategoryId = category.Id
	return nil
}
func (c *Category) UpdataCategory(ctx context.Context, r *pb.UpdataRequest, w *pb.UpdataResponse) error {
	category := &model.Category{
		CategoryName:        r.CategoryName,
		CategoryLevel:       r.CategoryLevel,
		CategoryParent:      r.CategoryParent,
		CategoryImage:       r.CategoryImage,
		CategoryDescription: r.CategoryDescription,
	}
	if err := c.CategoryDataService.UpdataCategory(category); err != nil {
		return err
	}
	w.Message = "更新成功"
	return nil
}
func (c *Category) DelectCategory(ctx context.Context, r *pb.DelectRequest, w *pb.DelectResponse) error {
	return c.CategoryDataService.DelectCategory(r.CategoryId)
}
func (c *Category) FindCategoryByName(ctx context.Context, r *pb.FindByNameRequest, w *pb.PropertyResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(r.CategoryName)
	if err != nil {
		return err
	}
	w = categoryResponse(category)
	return nil
}
func (c *Category) FindCategoryByID(ctx context.Context, r *pb.FindByIDRequest, w *pb.PropertyResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(r.CategoryId)
	if err != nil {
		return err
	}
	w = categoryResponse(category)
	return nil
}
func (c *Category) FindCategoryByLevel(ctx context.Context, r *pb.FindByLevelRequest, w *pb.FindAllResponse) error {
	category, err := c.CategoryDataService.FindCategoryByLevel(r.CategoryLevel)
	if err != nil {
		return err
	}
	for _, v := range category {
		w.Category = make([]*pb.PropertyResponse, 0)
		w.Category = append(w.Category, categoryResponse(*v))
	}
	return nil
}
func (c *Category) FindCategoryByParent(ctx context.Context, r *pb.FindByParentRequest, w *pb.FindAllResponse) error {
	category, err := c.CategoryDataService.FindCategoryByParent(r.CategoryParent)
	if err != nil {
		return err
	}
	for _, v := range category {
		w.Category = make([]*pb.PropertyResponse, 0)
		w.Category = append(w.Category, categoryResponse(*v))
	}
	return nil
}
func (c *Category) FindAllCategory(ctx context.Context, r *pb.FindAllRequest, w *pb.FindAllResponse) error {
	category, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	for _, v := range category {
		w.Category = make([]*pb.PropertyResponse, 0)
		w.Category = append(w.Category, categoryResponse(*v))
	}
	return nil
}

func categoryResponse(category model.Category) *pb.PropertyResponse {

	return &pb.PropertyResponse{
		CategoryId:          category.Id,
		CategoryName:        category.CategoryName,
		CategoryLevel:       category.CategoryLevel,
		CategoryParent:      category.CategoryParent,
		CategoryImage:       category.CategoryImage,
		CategoryDescription: category.CategoryDescription,
	}
}
