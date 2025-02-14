package db

import (
	"log"

	"github.com/patrickishaf/lema/models"
)

func FindPostById(id uint) models.Post {
	var post models.Post
	getDB().Where(&models.Post{ID: id}).First(&post)
	return post
}

func FindPostsByUser(userId uint) []models.Post {
	var posts []models.Post
	getDB().Where(&models.Post{UserID: userId}).Find(&posts).Order("id DESC")
	return posts
}

func FindPostsByUserPaginated(userId uint, limit int, offset int) ([]models.Post, error) {
	var posts []models.Post
	err := getDB().Limit(limit).Offset(offset).Order("id desc").Find(&posts).Error
	if err != nil {
		log.Printf("failed to find paginated users buy id %d", userId)
		return nil, err
	}
	return posts, nil
}

func FindPostCountByUser(userId uint) int64 {
	var count int64
	getDB().Model(&models.Post{}).Where(&models.Post{UserID: userId}).Count(&count)
	return count
}

func InsertPost(post *models.Post) (models.Post, error) {
	result := getDB().Create(&post)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return models.Post{
		ID:     post.ID,
		UserID: post.UserID,
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}

func DeletePost(id uint) {
	getDB().Delete(&models.Post{}, id)
}
