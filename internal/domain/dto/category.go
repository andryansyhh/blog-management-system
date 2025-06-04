package dto

import "blog-management-system/internal/domain/model"

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   int   `json:"id"`
	Name string `json:"name"`
}

func NewCategoryResponse(c model.Category) CategoryResponse {
	return CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}
