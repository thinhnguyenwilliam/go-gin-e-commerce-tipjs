package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/wire"
)

type UserRouter struct{}

// Register user routes
func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouterPublic := router.Group("/user")

	// Inject with Wire
	userHandler := wire.InitUserHandler()

	// Public endpoints
	userRouterPublic.GET("/register", userHandler.Register)

	userRouterPublic.POST("/otp", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "otp endpoint"})
	})

	// Private endpoints (you can add middleware here later)
	userRouterPrivate := router.Group("/user")
	userRouterPrivate.GET("/get-info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "get user info"})
	})
}
