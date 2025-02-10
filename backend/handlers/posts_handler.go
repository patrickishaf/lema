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
		c.IndentedJSON(http.StatusBadRequest, common.CreateErrorResponse("invalid id param"))
		return
	}

	posts := db.FindPostsByUser(uint(userId))
	c.IndentedJSON(http.StatusOK, models.ReversePosts(posts))
}

func createPost(c *gin.Context) {
	var reqBody dtos.CreatePostReqBody

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "invalid request body")
		return
	}

	errors := common.ValidateStruct(reqBody)
	if errors != nil {
		c.IndentedJSON(http.StatusBadRequest, errors)
		return
	}

	newPost, err := db.InsertPost(&models.Post{
		AuthorId: reqBody.AuthorId,
		Title:    reqBody.Title,
		Body:     reqBody.Body,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

func deletePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.CreateErrorResponse("invalid post id"))
		return
	}

	existingPost := db.FindPostById(uint(postId))
	if existingPost.ID == 0 {
		errMsg := fmt.Sprintf("post with id %d not found", postId)
		c.IndentedJSON(http.StatusNotFound, errMsg)
		return
	}

	db.DeletePost(uint(postId))
	successMsg := fmt.Sprintf("post with id %d deleted successfully", postId)
	c.IndentedJSON(http.StatusOK, successMsg)
}

func RegisterPostHandlers(router *gin.Engine) {
	router.GET("/posts", getPostsByUserId)
	router.POST("/posts", createPost)
	router.DELETE("/posts/:id", deletePost)
}
