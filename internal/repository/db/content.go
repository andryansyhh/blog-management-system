package repository

import (
	"blog-management-system/internal/domain/model"
	"time"

	"gorm.io/gorm"
)

type ContentRepository interface {
	GetAll() ([]model.Content, error)
	GetByID(id int) (*model.Content, error)
	Create(post *model.Content) error
	Update(content *model.Content) error
	Delete(id int) error
	WithTransaction(func(tx *gorm.DB) error) error
}

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) ContentRepository {
	return &contentRepository{db}
}

func (r *contentRepository) GetAll() ([]model.Content, error) {
	var posts []model.Content
	err := r.db.Preload("User").Preload("Category").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *contentRepository) GetByID(id int) (*model.Content, error) {
	var post model.Content
	err := r.db.Preload("User").Preload("Category").First(&post, id).Error
	return &post, err
}

func (r *contentRepository) Create(post *model.Content) error {
	return r.db.Create(post).Error
}

func (r *contentRepository) Update(post *model.Content) error {
	return r.db.Model(&model.Content{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"title":       post.Title,
		"content":     post.Content,
		"category_id": post.CategoryID,
		"updated_at":  time.Now(),
	}).Error
}

func (r *contentRepository) Delete(id int) error {
	return r.db.Delete(&model.Content{}, id).Error
}

func (r *contentRepository) WithTransaction(f func(tx *gorm.DB) error) error {
	return r.db.Transaction(f)
}
