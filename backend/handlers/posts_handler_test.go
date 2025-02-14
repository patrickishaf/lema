package handlers

import (
	"net/http"
	"testing"

	"github.com/patrickishaf/lema/models"
	"github.com/stretchr/testify/assert"
)

func TestGetPostsByUserId(t *testing.T) {
	writer := makeRequest("GET", "/v1/posts?userId=1", "")
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestCreatePost(t *testing.T) {
	postData := models.Post{
		UserID: 1,
		Title:  "A sample post title",
		Body:   "This is the body of a test post 1",
	}
	writer := makeRequest("POST", "/v1/posts", postData)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestDeletePost(t *testing.T) {
	postData := models.Post{
		UserID: 1,
		Title:  "A sample post title",
		Body:   "This is the body of a test post 1",
	}
	makeRequest("POST", "/v1/posts", postData)

	writer := makeRequest("DELETE", "/v1/posts/1", "")
	assert.Equal(t, http.StatusOK, writer.Code)
}
