package handler

import (
	"blog-management-system/internal/domain/dto"
	"blog-management-system/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUC usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUC usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUC}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	data, err := h.categoryUC.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	if err := h.categoryUC.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category created"})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid category id"})
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	if err := h.categoryUC.Update(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category updated"})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid category id"})
		return
	}

	if err := h.categoryUC.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
