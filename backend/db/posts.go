package db

import "github.com/patrickishaf/lema/models"

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
