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
	pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid page number", "failed to get users")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid page size", "failed to get users")
		return
	}

	offset := pageNumber * pageSize
	users := db.FindUsers(pageSize, offset)
	count := db.FindUserCount()
	totalPages := math.Ceil(float64(count) / float64(pageSize))

	data := map[string]any{
		"pageNumber": pageNumber,
		"pageSize":   pageSize,
		"totalPages": totalPages,
		"data":       users,
	}
	common.SendResponse(c, http.StatusOK, data, "users fetched sucessfully")
}

func getUserCount(c *gin.Context) {
	numberOfUsers := db.FindUserCount()
	common.SendResponse(c, http.StatusOK, numberOfUsers, "user count retrieved successfully")
}

func getUserById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid id param", "failed to get user")
		return
	}

	existingUser := db.FindUserById(uint(userID))
	c.IndentedJSON(http.StatusOK, existingUser)
	common.SendResponse(c, http.StatusOK, existingUser, "user fetched successfully")
}

func RegisterUserHandlers(router *gin.Engine) {
	router.GET("/users", getUsers)
	router.GET("/users/count", getUserCount)
	router.GET("/users/:id", getUserById)
}
