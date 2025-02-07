package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

func RateLimiterMiddleware(rate limiter.Rate) gin.HandlerFunc {

	store := memory.NewStore()
	instance := limiter.New(store, rate)

	return func(c *gin.Context) {
		key := c.ClientIP()

		context, err := instance.Get(c, key)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		if context.Reached {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}
