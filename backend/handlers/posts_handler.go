package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/db"
	dtos "github.com/patrickishaf/lema/dto"
	"github.com/patrickishaf/lema/models"
)

func getPostsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid ID param", "failed to get posts by user id")
		return
	}

	pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid page number", "failed to get users")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "11"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid page size", "failed to get users")
		return
	}

	offset := (pageNumber - 1) * pageSize
	posts, err := db.FindPostsByUserPaginated(uint(userId), pageSize, offset)
	if err != nil {
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to get posts by user")
		return
	}
	count := db.FindPostCountByUser(uint(userId))
	totalPages := math.Ceil(float64(count) / float64(pageSize))

	data := map[string]any{
		"pageNumber": pageNumber,
		"pageSize":   pageSize,
		"totalPages": totalPages,
		"data":       posts,
	}
	if pageNumber > int(totalPages) {
		data["pageNumber"] = totalPages
	}
	common.SendResponse(c, http.StatusOK, data, "posts fetched successfully")
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
