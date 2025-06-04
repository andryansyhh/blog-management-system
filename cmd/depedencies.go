package cmd

import (
	"blog-management-system/internal/handler"
	api "blog-management-system/internal/repository/api"
	repository "blog-management-system/internal/repository/db"
	routes "blog-management-system/internal/router"
	"blog-management-system/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type Dependency struct {
	Config *Config
}

func InitDependencies(r *gin.Engine) *Dependency {
	cfg, err := Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// External dependencies
	dbConn, err := NewClientDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	redisClient := NewClientRedis(cfg)

	// quoteApi
	quoteApi := api.NewQuoteRepository()

	// User module
	userRepository := repository.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// Category module
	categoryRepository := repository.NewCategoryRepository(dbConn)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	// Content module
	contentRepository := repository.NewContentRepository(dbConn)
	contentUsecase := usecase.NewContentUsecase(contentRepository, userRepository, categoryRepository, quoteApi, dbConn, redisClient)
	contentHandler := handler.NewContentHandler(contentUsecase)

	// Setup Gin router
	routes.SetupRoutes(r, userHandler, contentHandler, categoryHandler)

	return &Dependency{Config: cfg}
}
