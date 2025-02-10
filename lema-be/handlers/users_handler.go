package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/common"
	"github.com/patrickishaf/lema-be/db"
)

func getUsers(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := pageNumber * pageSize
	users := db.FindUsers(pageSize, offset)
	count := db.FindUserCount()
	totalPages := math.Ceil(float64(count) / float64(pageSize))

	c.IndentedJSON(http.StatusOK, map[string]any{
		"pageNumber": pageNumber,
		"pageSize":   pageSize,
		"totalPages": totalPages,
		"data":       users,
	})
}

func getUserCount(c *gin.Context) {
	numberOfusers := db.FindUserCount()
	c.IndentedJSON(http.StatusOK, numberOfusers)
}

func getUserById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.CreateErrorResponse("invalid id param"))
		return
	}

	existingUser := db.FindUserById(uint(userID))
	c.IndentedJSON(http.StatusOK, existingUser)
}

func RegisterUserHandlers(router *gin.Engine) {
	router.GET("/users", getUsers)
	router.GET("/users/count", getUserCount)
	router.GET("/users/:id", getUserById)
}
