package db

import "github.com/patrickishaf/lema-be/models"

func FindPostById(id uint) models.Post {
	var post models.Post
	getDB().Where(&models.Post{ID: id}).First(&post)
	return post
}

func FindPostsByUser(userId uint) []models.Post {
	var posts []models.Post
	getDB().Where(&models.Post{AuthorId: userId}).Find(&posts).Order("id DESC")
	return posts
}

func InsertPost(post *models.Post) (models.Post, error) {
	result := getDB().Create(&post)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return models.Post{
		ID:       post.ID,
		AuthorId: post.AuthorId,
		Title:    post.Title,
		Body:     post.Body,
	}, nil
}

func DeletePost(id uint) {
	getDB().Delete(&models.Post{}, id)
}

func insertDummyPosts() error {
	anchorText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dol..."
	posts := []models.Post{
		{
			AuthorId: 1,
			Title:    "I got a letter",
			Body:     anchorText,
		},
		{
			AuthorId: 1,
			Title:    "What a nice town",
			Body:     anchorText,
		},
		{
			AuthorId: 1,
			Title:    "I love my wife, Mary",
			Body:     anchorText,
		},
		{
			AuthorId: 1,
			Title:    "How can anyone eat pizza at a time like this?",
			Body:     anchorText,
		},
		{
			AuthorId: 2,
			Title:    "I love my husband, Brad",
			Body:     anchorText,
		},
		{
			AuthorId: 2,
			Title:    "I can definitely eat pizza at a time like this",
			Body:     anchorText,
		},
	}
	result := getDB().Create(&posts)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
