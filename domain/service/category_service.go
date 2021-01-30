package service

import (
	"github.com/itgamerskyHuang/category/domain/model"
	"github.com/itgamerskyHuang/category/domain/repository"
	"github.com/jinzhu/gorm"
)

type ICategoryService interface {
	CreateCategory(*model.Category) error
	UpdataCategory(*model.Category) error
	DelectCategory(int64) error
	FindCategoryByName(string) (model.Category, error)
	FindCategoryByID(int64) (model.Category, error)
	FindCategoryByLevel(uint32) ([]*model.Category, error)
	FindCategoryByParent(int64) ([]*model.Category, error)
	FindAllCategory() ([]*model.Category, error)
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

func NewCategoryService(db *gorm.DB) ICategoryService {
	return &CategoryService{
		CategoryRepository: repository.NewCategoryRepository(db),
	}
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
	return s.CreateCategory(category)
}
func (s *CategoryService) UpdataCategory(category *model.Category) error {
	return s.UpdataCategory(category)
}
func (s *CategoryService) DelectCategory(id int64) error {
	return s.CategoryRepository.DelectCategory(id)
}
func (s *CategoryService) FindCategoryByName(name string) (model.Category, error) {
	return s.FindCategoryByName(name)
}
func (s *CategoryService) FindCategoryByID(id int64) (model.Category, error) {
	return s.CategoryRepository.FindCategoryByID(id)
}
func (s *CategoryService) FindCategoryByLevel(level uint32) ([]*model.Category, error) {
	return s.CategoryRepository.FindCategoryByLevel(level)
}
func (s *CategoryService) FindCategoryByParent(parent int64) ([]*model.Category, error) {
	return s.CategoryRepository.FindCategoryByParent(parent)
}
func (s *CategoryService) FindAllCategory() ([]*model.Category, error) {
	return s.CategoryRepository.FindAllCategory()
}
