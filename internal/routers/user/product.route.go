package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

// Register product routes
func (pr *ProductRouter) InitProductRouter(router *gin.RouterGroup) {
	productRouterPublic := router.Group("/product")

	// Define public endpoints
	productRouterPublic.GET("/search", func(c *gin.Context) {
		// placeholder logic
		c.JSON(http.StatusOK, gin.H{"message": "Product search"})
	})

	productRouterPublic.GET("/detail/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "Product detail", "id": id})
	})
}
