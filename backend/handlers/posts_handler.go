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
	userId := c.DefaultQuery("userId", "...")
	if userId == "..." {
		log.Printf("failed to get posts by user id. id not provided")
		common.SendResponse(c, http.StatusBadRequest, "invalid ID param", "failed to get posts by user id")
		return
	}
	pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	if err != nil {
		log.Printf("failed to get posts by user id %s. error: %v", userId, err)
		common.SendResponse(c, http.StatusBadRequest, "invalid page number", "failed to get users")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "11"))
	if err != nil {
		log.Printf("failed to get posts by user id %s. error: %v", userId, err)
		common.SendResponse(c, http.StatusBadRequest, "invalid page size", "failed to get users")
		return
	}

	postsPaginated, err := h.repo.FindPostsByUserID(userId, pageNumber, pageSize)
	if err != nil {
		log.Printf("failed to get posts with user id %s. error: %v", userId, err)
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

	newPost, err := h.repo.InsertPost(reqBody.Title, reqBody.Body, reqBody.UserID)
	if err != nil {
		log.Printf("failed to create post. error: %v", err)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to create post")
		return
	}

	common.SendResponse(c, http.StatusCreated, newPost, "post created succesfully")
}

func (h *PostsHandler) deletePost(c *gin.Context) {
	// TODO: Add field validation here
	postId := c.Param("id")

	existingPost, err := h.repo.FindPostById(postId)
	if err != nil {
		errMsg := fmt.Sprintf("post with id %s not found", postId)
		common.SendResponse(c, http.StatusNotFound, []string{}, errMsg)
		return
	}

	err = h.repo.DeletePost(postId)
	if err != nil {
		log.Printf("failed to delete post with id %s", postId)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to delete post")
		return
	}
	successMsg := fmt.Sprintf("post with id %s deleted successfully", postId)
	common.SendResponse(c, http.StatusOK, existingPost, successMsg)
}

func (h *PostsHandler) RegisterRequestHandlers(router *gin.RouterGroup) {
	router.GET("/posts", h.getPostsByUserId)
	router.POST("/posts", h.createPost)
	router.DELETE("/posts/:id", h.deletePost)
}
