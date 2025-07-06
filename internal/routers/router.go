package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/thinhcompany/ecommerce-ver-2/internal/controller"
	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
)

func SetupRoutes(r *gin.Engine) {
	// Step 1: Setup dependencies manually
	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := c.NewUserController(userService)
	pongController := c.NewPongController() // if your PongController has no dependencies

	// Group routes under /v1/2024
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", pongController.PingHandler)
		v1.POST("/hello", HelloHandler)
		v1.GET("/hello/:name", userController.HelloByNameHandler)
		v1.GET("/users", userController.GetUserInfoHandler)
	}
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
