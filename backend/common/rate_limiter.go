package common

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limiters = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func PlaceRateLimits() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		limiter, exists := limiters[ip]
		if !exists {
			limiter = rate.NewLimiter(1, 5)
			limiters[ip] = limiter
		}
		mu.Unlock()

		if !limiter.Allow() {
			SendResponse(c, http.StatusTooManyRequests, "too many requests", "failed to execute request")
			c.Abort()
			return
		}
		c.Next()
	}
}
