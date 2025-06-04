package routes

import (
	"blog-management-system/internal/handler"
	"blog-management-system/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userHandler *handler.UserHandler,
	contentHandler *handler.ContentHandler,
	categoryHandler *handler.CategoryHandler,
) {
	api := router.Group("/v1")
	useMiddleware := middleware.AuthMiddleware()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	// User (Auth)
	userGroup := api.Group("/user")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
	}

	// Content (Post)
	contentGroup := api.Group("/content")
	{
		contentGroup.GET("/", contentHandler.GetAllPosts)
		contentGroup.GET("/:id", contentHandler.GetPostByID)
		contentGroup.POST("/", useMiddleware, contentHandler.CreatePost)
		contentGroup.PUT("/:id", useMiddleware, contentHandler.UpdatePost)
		contentGroup.DELETE("/:id", useMiddleware, contentHandler.DeletePost)
	}

	// Category
	categoryGroup := api.Group("/category").Use(useMiddleware)
	{
		categoryGroup.GET("/", categoryHandler.GetAll)
		categoryGroup.POST("/", categoryHandler.Create)
		categoryGroup.PUT("/:id", categoryHandler.Update)
	}

	// 💡 Optional: Public API
	api.GET("/quote", contentHandler.GetQuote)
}
