package db

import (
	"log"
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

func (r *PostsRepository) FindPostsByUserID(userID string) (*[]models.Post, error) {
	var posts []models.Post
	err := r.db.Order("updated_at desc").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		log.Printf("error in PostsRepository.FindPostsByUserID: %v", err)
		return nil, err
	}

	return &posts, nil
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
