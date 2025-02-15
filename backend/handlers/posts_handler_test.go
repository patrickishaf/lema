package handlers

import (
	"net/http"
	"testing"

	"github.com/patrickishaf/lema/db"
	"github.com/patrickishaf/lema/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
	db.PostsRepository
}

func (m *MockRepository) FindPostsByUserID(userID uint, pageNumber, pageSize int) (interface{}, error) {
	args := m.Called(userID, pageNumber, pageSize)
	return args.Get(0), args.Error(1)
}

func TestGetPostsByUserId(t *testing.T) {
	writer := makeRequest("GET", "/v1/posts?userId=1", "")
	assert.Equal(t, http.StatusOK, writer.Code)

	writer = makeRequest("GET", "/v1/posts?userId=30", "")
	assert.Equal(t, http.StatusInternalServerError, writer.Code)
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
