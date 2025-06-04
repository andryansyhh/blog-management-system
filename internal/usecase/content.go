package usecase

import (
	"blog-management-system/internal/domain/dto"
	"blog-management-system/internal/domain/model"
	api "blog-management-system/internal/repository/api"
	repository "blog-management-system/internal/repository/db"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ContentUsecase interface {
	GetAllPosts() ([]dto.ContentResponse, error)
	GetPostByID(id int) (*dto.ContentResponse, error)
	CreatePost(req dto.CreateContentRequest, userID int) error
	UpdatePost(postID int, userID int, req dto.UpdateContentRequest) error
	DeletePost(postID int, userID int) error
	GetQuote() (string, string, string, error)
}

type contentUsecase struct {
	contentRepo repository.ContentRepository
	userRepo    repository.UserRepository
	catRepo     repository.CategoryRepository
	quoteRepo   api.QuoteRepository
	db          *gorm.DB
	redis       *redis.Client
}

func NewContentUsecase(contentRepo repository.ContentRepository, userRepo repository.UserRepository, catRepo repository.CategoryRepository, quoteRepo api.QuoteRepository, db *gorm.DB, redis *redis.Client) ContentUsecase {
	return &contentUsecase{contentRepo, userRepo, catRepo, quoteRepo, db, redis}
}

func (uc *contentUsecase) GetAllPosts() ([]dto.ContentResponse, error) {
	ctx := context.Background()
	cacheKey := "content:all"

	// redis GetAllPosts
	cached, _ := uc.redis.Get(ctx, cacheKey).Result()
	if cached != "" {
		var data []dto.ContentResponse
		if json.Unmarshal([]byte(cached), &data) == nil {
			if len(data) > 0 {
				return data, nil // return jika cache valid dan tidak kosong
			}
		}
	}

	// Fallback to DB
	posts, err := uc.contentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.ContentResponse
	for _, p := range posts {
		result = append(result, dto.NewPostResponse(p))
	}

	// Fetch from DB and cache only if not empty
	if len(result) > 0 {
		data, _ := json.Marshal(result)
		uc.redis.Set(ctx, cacheKey, data, 2*time.Minute)
	}

	return result, nil
}

func (uc *contentUsecase) GetPostByID(id int) (*dto.ContentResponse, error) {
	post, err := uc.contentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := dto.NewPostResponse(*post)
	return &resp, nil
}

func (uc *contentUsecase) CreatePost(req dto.CreateContentRequest, userID int) error {
	err := uc.db.Transaction(func(tx *gorm.DB) error {
		// Init transactional repos
		userRepoTx := repository.NewUserRepository(tx)
		catRepoTx := repository.NewCategoryRepository(tx)
		contentRepoTx := repository.NewContentRepository(tx)

		// 1. Validasi user
		user, err := userRepoTx.GetById(userID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		// 2. Validasi category
		category, err := catRepoTx.GetByID(req.CategoryID)
		if err != nil {
			return fmt.Errorf("category not found: %w", err)
		}

		// 3. Buat post
		post := model.Content{
			Title:      req.Title,
			Content:    req.Content,
			UserID:     user.ID,
			CategoryID: category.ID,
		}

		err = contentRepoTx.Create(&post)
		if err != nil {
			return fmt.Errorf("failed to create post: %w", err)
		}

		return nil
	})

	return err
}

func (uc *contentUsecase) UpdatePost(postID int, userID int, req dto.UpdateContentRequest) error {
	post, err := uc.contentRepo.GetByID(postID)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return errors.New("unauthorized: cannot edit this post")
	}

	post.Title = req.Title
	post.Content = req.Content
	post.CategoryID = req.CategoryID

	log.Println(post.CategoryID)

	if err := uc.contentRepo.Update(post); err != nil {
		return err
	}

	uc.redis.Del(context.Background(), "content:all") // Invalidate cache
	return nil
}

func (uc *contentUsecase) DeletePost(postID int, userID int) error {
	post, err := uc.contentRepo.GetByID(postID)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return errors.New("unauthorized: cannot delete this post")
	}

	if err := uc.contentRepo.Delete(postID); err != nil {
		return err
	}

	uc.redis.Del(context.Background(), "content:all") // Invalidate cache
	return nil
}

func (uc *contentUsecase) GetQuote() (string, string, string, error) {
	return uc.quoteRepo.GetRandomQuote()
}
