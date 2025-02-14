package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/db"
	dtos "github.com/patrickishaf/lema/dto"
	"github.com/patrickishaf/lema/models"
)

type PostsHandler struct {
	repo *db.PostsRepository
}

func NewPostsHandler(repo *db.PostsRepository) *PostsHandler {
	return &PostsHandler{
		repo: repo,
	}
}

func (h *PostsHandler) getPostsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		log.Printf("failed to get posts by user id. error: %v", err)
		common.SendResponse(c, http.StatusBadRequest, "invalid ID param", "failed to get posts by user id")
		return
	}

	pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	if err != nil {
		log.Printf("failed to get posts by user id %d. error: %v", userId, err)
		common.SendResponse(c, http.StatusBadRequest, "invalid page number", "failed to get users")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "11"))
	if err != nil {
		log.Printf("failed to get posts by user id %d. error: %v", userId, err)
		common.SendResponse(c, http.StatusBadRequest, "invalid page size", "failed to get users")
		return
	}

	postsPaginated, err := h.repo.FindPostsByUserID(uint(userId), pageNumber, pageSize)
	if err != nil {
		log.Printf("failed to get posts with user id %d. error: %v", userId, err)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to get posts by user")
		return
	}
	common.SendResponse(c, http.StatusOK, postsPaginated, "posts fetched successfully")
}

func (h *PostsHandler) createPost(c *gin.Context) {
	var reqBody dtos.CreatePostReqBody

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Printf("failed to create post. error: %v", err)
		common.SendResponse(c, http.StatusBadRequest, "invalid request body", "failed to create post")
		return
	}

	errors := common.ValidateStruct(reqBody)
	if errors != nil {
		log.Printf("failed to create post. errors: %v", errors)
		common.SendResponse(c, http.StatusBadRequest, errors, "failed to create post")
		return
	}

	newPost, err := h.repo.InsertPost(&models.Post{
		UserID: reqBody.UserID,
		Title:  reqBody.Title,
		Body:   reqBody.Body,
	})
	if err != nil {
		log.Printf("failed to create post. error: %v", err)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to create post")
		return
	}

	common.SendResponse(c, http.StatusCreated, newPost, "post created succesfully")
}

func (h *PostsHandler) deletePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("failed to delete post. error: %v", err)
		common.SendResponse(c, http.StatusBadRequest, []string{}, "invalid post id")
		return
	}

	existingPost, err := h.repo.FindPostById(uint(postId))
	if err != nil {
		errMsg := fmt.Sprintf("post with id %d not found", postId)
		common.SendResponse(c, http.StatusNotFound, []string{}, errMsg)
		return
	}

	err = h.repo.DeletePost(uint(postId))
	if err != nil {
		log.Printf("failed to delete post with id %d", postId)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to delete post")
		return
	}
	successMsg := fmt.Sprintf("post with id %d deleted successfully", postId)
	common.SendResponse(c, http.StatusOK, existingPost, successMsg)
}

func (h *PostsHandler) RegisterRequestHandlers(router *gin.RouterGroup) {
	router.GET("/posts", h.getPostsByUserId)
	router.POST("/posts", h.createPost)
	router.DELETE("/posts/:id", h.deletePost)
}
