package handler

import (
	"blog-management-system/internal/domain/dto"
	"blog-management-system/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	contentUC usecase.ContentUsecase
}

func NewContentHandler(contentUC usecase.ContentUsecase) *ContentHandler {
	return &ContentHandler{contentUC}
}

func (h *ContentHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.contentUC.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *ContentHandler) GetPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.contentUC.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func (h *ContentHandler) CreatePost(c *gin.Context) {
	var req dto.CreateContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	// Ambil user ID dari JWT claims (assume middleware set "user_id")
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	userID := userIDRaw.(int)

	if err := h.contentUC.CreatePost(req, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created"})
}

func (h *ContentHandler) UpdatePost(c *gin.Context) {
	var req dto.UpdateContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	postID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")

	if err := h.contentUC.UpdatePost(postID, userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post updated"})
}

func (h *ContentHandler) DeletePost(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")

	if err := h.contentUC.DeletePost(postID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

func (h *ContentHandler) GetQuote(c *gin.Context) {
	quote, author, blockQuote, err := h.contentUC.GetQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get quote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quote":      quote,
		"author":     author,
		"blockquote": blockQuote,
	})
}
