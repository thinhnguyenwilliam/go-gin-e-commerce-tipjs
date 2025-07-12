package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// Register user routes
func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouterPublic := router.Group("/user")

	// Public endpoints
	userRouterPublic.GET("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "register endpoint"})
	})

	userRouterPublic.POST("/otp", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "otp endpoint"})
	})

	// Private endpoints (you can add middleware here later)
	userRouterPrivate := router.Group("/user")
	userRouterPrivate.GET("/get-info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "get user info"})
	})
}
