package middlewares

import (
	"net/http"
	"time"

	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"go.uber.org/zap"
)

func RateLimiterMiddleware(maxRequests int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rl:%s", ip) // rate limit key per IP
		ctx := context.Background()

		// Increment request count
		count, err := global.Rdb.Incr(ctx, key).Result()
		if err != nil {
			global.AppLogger.Error("Redis INCR failed", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Rate limit error"})
			return
		}

		if count == 1 {
			// First time, set expiration
			global.Rdb.Expire(ctx, key, window)
		}

		if int(count) > maxRequests {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}

		c.Next()
	}
}
