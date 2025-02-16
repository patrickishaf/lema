package db

import (
	"log"
	"math"
	"time"

	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/models"
	"gorm.io/gorm"
)

type PostsRepository struct {
	db *gorm.DB
}

func NewPostsRepository() *PostsRepository {
	return &PostsRepository{
		db: getDB(),
	}
}

func (r *PostsRepository) FindPostById(id string) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where(&models.Post{ID: id}).First(&post).Error; err != nil {
		log.Printf("error in PostsRepository.FindPostByID: %v", err)
		return nil, err
	}
	return &post, nil
}

func (r *PostsRepository) FindPostsByUserID(userID string, pageNumber int, limit int) (*common.PaginatedItems[models.Post], error) {
	var posts []models.Post

	totalPostsByUser, err := r.findPostCountByUser(userID)
	if err != nil {
		return nil, err
	}

	if totalPostsByUser == 0 {
		log.Println("user has no posts")
		return &common.PaginatedItems[models.Post]{
			Data:       []models.Post{},
			PageNumber: 0,
			PageSize:   limit,
			TotalPages: 0,
		}, nil
	}

	// get the elements at the final page if the pageNumber is out of bounds
	totalPages := math.Ceil(float64(totalPostsByUser) / float64(limit))
	if pageNumber > int(totalPages) {
		pageNumber = int(totalPages)
	}

	offset := (pageNumber - 1) * limit
	err = r.db.Limit(limit).Offset(offset).Order("updated_at desc").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		log.Printf("error in PostsRepository.FindPostsByUserID: %v", err)
		return nil, err
	}

	return &common.PaginatedItems[models.Post]{
		Data:       posts,
		PageNumber: pageNumber,
		PageSize:   limit,
		TotalPages: int(totalPages),
		TotalItems: int(totalPostsByUser),
	}, nil
}

func (r *PostsRepository) findPostCountByUser(userID string) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Post{}).Where(&models.Post{UserID: userID}).Count(&count).Error; err != nil {
		log.Printf("failed to get post count by user id %v", err)
		return 0, err
	}
	return count, nil
}

func (r *PostsRepository) InsertPost(title string, body string, userId string) (*models.Post, error) {
	id, _ := common.GenerateID()
	now := time.Now()
	createdAt := now.Format(time.RFC3339)
	post := models.Post{
		ID:        id,
		UserID:    userId,
		Title:     title,
		Body:      body,
		CreatedAt: createdAt,
	}
	result := r.db.Create(&post)
	if result.Error != nil {
		log.Printf("error in PostsRepository.InsertPost: %v", result.Error)
		return &models.Post{}, result.Error
	}

	return &models.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
	}, nil
}

func (r *PostsRepository) DeletePost(postID string) error {
	var post models.Post
	err := r.db.Where("id = ?", postID).Delete(&post).Error
	if err != nil {
		log.Printf("error in PostsRepository.DeletePost: %v", err)
		return err
	}
	return nil
}
