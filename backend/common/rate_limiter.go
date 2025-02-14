package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 5)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			SendResponse(c, http.StatusTooManyRequests, "too many requests", "failed to execute request")
			c.Abort()
			return
		}
		c.Next()
	}
}
