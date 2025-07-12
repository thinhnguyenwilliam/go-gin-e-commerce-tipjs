package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

// Register admin routes
func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	adminRouterPublic := router.Group("/admin")

	// Public login endpoint
	adminRouterPublic.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "admin login"})
	})

	// Private endpoints (e.g., require authentication middleware)
	adminRouterPrivate := router.Group("/admin/user")

	adminRouterPrivate.POST("/active-user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "user activated"})
	})
}
