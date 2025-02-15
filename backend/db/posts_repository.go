package db

import (
	"log"
	"math"

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
	err = r.db.Limit(limit).Offset(offset).Order("id desc").Where("user_id = ?", userID).Find(&posts).Error
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

func (r *PostsRepository) InsertPost(post *models.Post) (*models.Post, error) {
	result := r.db.Create(&post)
	if result.Error != nil {
		log.Printf("error in PostsRepository.InsertPost: %v", result.Error)
		return &models.Post{}, result.Error
	}

	return &models.Post{
		ID:     post.ID,
		UserID: post.UserID,
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}

func (r *PostsRepository) DeletePost(postID string) error {
	err := r.db.Delete(&models.Post{}, postID).Error
	if err != nil {
		log.Printf("error in PostsRepository.DeletePost: %v", err)
		return err
	}
	return nil
}
