package dto

import (
	"blog-management-system/internal/domain/model"
	"time"
)

type CreateContentRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type UpdateContentRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type ContentResponse struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Author    SimpleUserDTO `json:"author"`
	Category  string        `json:"category"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}

type SimpleUserDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewPostResponse(content model.Content) ContentResponse {
	return ContentResponse{
		ID:      content.ID,
		Title:   content.Title,
		Content: content.Content,
		Author: SimpleUserDTO{
			ID:   content.User.ID,
			Name: content.User.Name,
		},
		Category:  content.Category.Name,
		CreatedAt: content.CreatedAt.Format(time.RFC3339),
		UpdatedAt: content.UpdatedAt.Format(time.RFC3339),
	}
}
