package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/controller"
)

func SetupRoutes(r *gin.Engine) {
	// Group routes under /v1/2024
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", controller.PingHandler)
		v1.POST("/hello", HelloHandler)
		v1.GET("/hello/:name", HelloByNameHandler)
	}
}

func HelloByNameHandler(c *gin.Context) {
	name := c.Param("name")              // e.g. /hello/thinh
	uid := c.DefaultQuery("uid", "0000") // ?uid=123 or default "0000"

	log.Println("Name:", name, "UID:", uid)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "thinh"},
	})
}

func HelloHandler(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hello " + body.Name})
}
