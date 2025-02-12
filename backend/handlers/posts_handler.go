package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/common"
	"github.com/patrickishaf/lema-be/db"
	dtos "github.com/patrickishaf/lema-be/dto"
	"github.com/patrickishaf/lema-be/models"
)

func getPostsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid ID param", "failed to get posts by user id")
		return
	}

	posts := db.FindPostsByUser(uint(userId))
	common.SendResponse(c, http.StatusOK, models.ReversePosts(posts), "posts fetched successfully")
}

func createPost(c *gin.Context) {
	var reqBody dtos.CreatePostReqBody

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid request body", "failed to create post")
		return
	}

	errors := common.ValidateStruct(reqBody)
	if errors != nil {
		common.SendResponse(c, http.StatusBadRequest, errors, "failed to create post")
		return
	}

	newPost, err := db.InsertPost(&models.Post{
		UserID: reqBody.UserID,
		Title:  reqBody.Title,
		Body:   reqBody.Body,
	})
	if err != nil {
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to create post")
		return
	}

	common.SendResponse(c, http.StatusCreated, newPost, "post created succesfully")
}

func deletePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, []string{}, "invalid post id")
		return
	}

	existingPost := db.FindPostById(uint(postId))
	if existingPost.ID == 0 {
		errMsg := fmt.Sprintf("post with id %d not found", postId)
		common.SendResponse(c, http.StatusNotFound, []string{}, errMsg)
		return
	}

	db.DeletePost(uint(postId))
	successMsg := fmt.Sprintf("post with id %d deleted successfully", postId)
	common.SendResponse(c, http.StatusOK, existingPost, successMsg)
}

func RegisterPostHandlers(router *gin.Engine) {
	router.GET("/posts", getPostsByUserId)
	router.POST("/posts", createPost)
	router.DELETE("/posts/:id", deletePost)
}
