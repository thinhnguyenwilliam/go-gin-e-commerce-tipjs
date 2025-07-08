package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Println("Raw Authorization header:", authHeader) // 👈 log before processing
		const prefix = "Bearer "

		if !strings.HasPrefix(authHeader, prefix) {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse(response.ErrorCodeTokenInvalid, nil))
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, prefix)

		// ✅ Log token to console
		log.Println("Extracted Token:", token)

		// TODO: Validate the token (e.g., parse JWT, check expiration, etc.)
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse(response.ErrorCodeTokenInvalid, nil))
			c.Abort()
			return
		}

		// Token is valid (assuming validation logic passes)
		c.Next()
	}
}
