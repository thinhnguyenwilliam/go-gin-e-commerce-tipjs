package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouterPrivate := router.Group("/admin/user")

	userRouterPrivate.POST("/active-user-1", func(c *gin.Context) {
		// Example logic: activate user
		c.JSON(http.StatusOK, gin.H{
			"message": "User activated",
		})
	})
}
