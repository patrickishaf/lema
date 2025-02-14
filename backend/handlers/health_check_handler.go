package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/db"
)

type HealthCheckHandler struct {
	repo *db.HealthCheckRepository
}

func NewHealthCheckHandler(repo *db.HealthCheckRepository) *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) healthCheck(c *gin.Context) {
	isHealthy := h.repo.IsDatabaseOnline()
	if isHealthy {
		common.SendResponse(c, http.StatusOK, "server is healthy", "health check passed")
	} else {
		common.SendResponse(c, http.StatusServiceUnavailable, "database connectionn failed", "health check failed")
	}
}

func (h *HealthCheckHandler) RegisterRequestHandlers(router *gin.RouterGroup) {
	router.GET("/health", h.healthCheck)
}
