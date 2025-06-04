package repository

import (
	"blog-management-system/internal/domain/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetByID(id int) (*model.Category, error)
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id int) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int) error {
	return r.db.Delete(&model.Category{}, id).Error
}
