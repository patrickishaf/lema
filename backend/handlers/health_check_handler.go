package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
)

func healthCheck(c *gin.Context) {
	common.SendResponse(c, http.StatusOK, "server is healthy", "health check passed")
}

func RegisterHealthCheckHandlers(router *gin.Engine) {
	router.GET("/health", healthCheck)
}
