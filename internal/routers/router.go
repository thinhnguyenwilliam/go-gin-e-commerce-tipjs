package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/thinhcompany/ecommerce-ver-2/internal/controller"
	"github.com/thinhcompany/ecommerce-ver-2/internal/middlewares"
	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
)

func SetupRoutes(r *gin.Engine) {
	// Step 1: Setup dependencies manually
	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := c.NewUserController(userService)
	pongController := c.NewPongController() // if your PongController has no dependencies

	// Global middleware: Rate limiting
	r.Use(middlewares.RateLimitMiddleware())

	// Group routes under /v1/2024
	// Public routes
	public := r.Group("/v1/2024")
	{
		public.GET("/ping", pongController.PingHandler)
		public.POST("/hello", HelloHandler)
	}

	// Protected routes
	protected := r.Group("/v1/2024", middlewares.AuthenMiddleware())
	{
		protected.GET("/hello/:name", userController.HelloByNameHandler)
		protected.GET("/users", userController.GetUserInfoHandler)
		protected.GET("/user/:id", userController.GetUserByID)
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
