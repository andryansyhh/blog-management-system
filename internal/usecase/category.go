package usecase

import (
	"blog-management-system/internal/domain/dto"
	"blog-management-system/internal/domain/model"
	repository "blog-management-system/internal/repository/db"
	"fmt"
)

type CategoryUsecase interface {
	GetAll() ([]dto.CategoryResponse, error)
	Create(req dto.CreateCategoryRequest) error
	Update(id int, req dto.UpdateCategoryRequest) error
	Delete(id int) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo}
}

func (uc *categoryUsecase) GetAll() ([]dto.CategoryResponse, error) {
	data, err := uc.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.CategoryResponse
	for _, c := range data {
		result = append(result, dto.NewCategoryResponse(c))
	}
	return result, nil
}

func (uc *categoryUsecase) Create(req dto.CreateCategoryRequest) error {
	category := model.Category{
		Name: req.Name,
	}
	return uc.categoryRepo.Create(&category)
}

func (uc *categoryUsecase) Update(id int, req dto.UpdateCategoryRequest) error {
	category, err := uc.categoryRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	category.Name = req.Name
	return uc.categoryRepo.Update(category)
}

func (uc *categoryUsecase) Delete(id int) error {
	return uc.categoryRepo.Delete(id)
}
