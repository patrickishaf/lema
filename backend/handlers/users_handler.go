package handlers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/db"
)

type UsersHandler struct {
	repo *db.UserRepository
}

func NewUsersHandler(repo *db.UserRepository) *UsersHandler {
	return &UsersHandler{
		repo: repo,
	}
}

func (h *UsersHandler) getUsers(c *gin.Context) {
	pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid page number", "failed to get users")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	if err != nil {
		common.SendResponse(c, http.StatusBadRequest, "invalid page size", "failed to get users")
		return
	}

	users, err := h.repo.FindAll(pageNumber, pageSize)
	if err != nil {
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to get users")
		return
	}
	count, err := h.repo.GetUserCount()
	if err != nil {
		log.Printf("failed to get user count. error: %v", err)
		common.SendResponse(c, http.StatusInternalServerError, err, "fialed to get user count")
		return
	}
	totalPages := math.Ceil(float64(count) / float64(pageSize))

	data := map[string]any{
		"pageNumber": pageNumber,
		"pageSize":   pageSize,
		"totalPages": totalPages,
		"data":       users,
	}
	if pageNumber > int(totalPages) {
		data["pageNumber"] = totalPages
	}
	common.SendResponse(c, http.StatusOK, data, "users fetched sucessfully")
}

func (h *UsersHandler) getUserCount(c *gin.Context) {
	numberOfUsers, err := h.repo.GetUserCount()
	if err != nil {
		fmt.Printf("failed to get user count. error: %v", err)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to get user count")
		return
	}
	common.SendResponse(c, http.StatusOK, numberOfUsers, "user count retrieved successfully")
}

func (h *UsersHandler) getUserById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to get user by ID. error: %v", err)
		common.SendResponse(c, http.StatusBadRequest, "invalid id param", "failed to get user")
		return
	}

	existingUser, err := h.repo.FindUserById(uint(userID))
	if err != nil {
		fmt.Printf("failed to get user by ID. error: %v", err)
		common.SendResponse(c, http.StatusInternalServerError, err, "failed to get user by ID")
		return
	}
	common.SendResponse(c, http.StatusOK, existingUser, "user fetched successfully")
}

func (h *UsersHandler) RegisterRequestHandlers(router *gin.Engine) {
	router.GET("/users", h.getUsers)
	router.GET("/users/count", h.getUserCount)
	router.GET("/users/:id", h.getUserById)
}
