package db

import (
	"log"

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

func (r *PostsRepository) FindPostById(id uint) (*models.Post, error) {
	return &models.Post{}, nil
}

func (r *PostsRepository) FindPostsByUserID(userID uint, limit int, offset int) (*[]models.Post, error) {
	var posts []models.Post
	err := r.db.Limit(limit).Offset(offset).Order("id desc").Find(&posts).Error
	if err != nil {
		log.Printf("error in PostsRepository.FindPostsByUserID: %v", err)
		return nil, err
	}
	return &posts, nil
}

func (r *PostsRepository) findPostCountByUser(userID uint) int64 {
	return 0
}

func (r *PostsRepository) InsertPost(post *models.Post) (*models.Post, error) {
	return &models.Post{}, nil
}

func (r *PostsRepository) DeletePost(postID uint) error {
	return nil
}
