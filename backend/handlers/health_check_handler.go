package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) healthCheck(c *gin.Context) {
	common.SendResponse(c, http.StatusOK, "server is healthy", "health check passed")
}

func (h *HealthCheckHandler) RegisterRequestHandlers(router *gin.Engine) {
	router.GET("/health", h.healthCheck)
}
