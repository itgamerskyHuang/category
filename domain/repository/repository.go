package repository

import (
	"github.com/itgamerskyHuang/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	CreateCategory(*model.Category) error
	UpdataCategory(*model.Category) error
	DelectCategory(int64) error
	FindCategoryByName(string) (model.Category, error)
	FindCategoryByID(int64) (model.Category, error)
	FindCategoryByLevel(uint32) ([]*model.Category, error)
	FindCategoryByParent(int64) ([]*model.Category, error)
	FindAllCategory() ([]*model.Category, error)
}

type CategoryRepository struct {
	mysql *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		mysql: db,
	}
}

func (c *CategoryRepository) InitTable() error {
	return c.mysql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Category{}).Error
}

func (c *CategoryRepository) CreateCategory(category *model.Category) error {
	return c.mysql.Create(category).Error
}

func (c *CategoryRepository) UpdataCategory(category *model.Category) error {
	return c.mysql.Model(category).Update(category).Error
}

func (c *CategoryRepository) DelectCategory(id int64) error {
	return c.mysql.Where("category_id = ?", id).Delete(&model.Category{}).Error
}

func (c *CategoryRepository) FindCategoryByName(name string) (category model.Category, err error) {
	return category, c.mysql.Where("category_name = ?", name).Find(&category).Error
}
func (c *CategoryRepository) FindCategoryByID(id int64) (category model.Category, err error) {
	return category, c.mysql.Where("category_id = ?", id).Find(&category).Error
}

func (c *CategoryRepository) FindCategoryByLevel(level uint32) (categorys []*model.Category, err error) {
	categorys = make([]*model.Category, 0)
	return categorys, c.mysql.Where("category_level = ?", level).Find(categorys).Error
}
func (c *CategoryRepository) FindCategoryByParent(parent int64) (categorys []*model.Category, err error) {
	categorys = make([]*model.Category, 0)
	return categorys, c.mysql.Where("category_parent = ?", parent).Find(categorys).Error
}
func (c *CategoryRepository) FindAllCategory() (categorys []*model.Category, err error) {
	categorys = make([]*model.Category, 0)
	return categorys, c.mysql.Find(categorys).Error
}
