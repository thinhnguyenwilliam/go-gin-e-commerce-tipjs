package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/thinhcompany/ecommerce-ver-2/internal/controller"
	"github.com/thinhcompany/ecommerce-ver-2/internal/middlewares"
	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
)

func InitRouter() *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// Step 1: Setup dependencies manually
	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := c.NewUserController(userService)
	pongController := c.NewPongController() // if your PongController has no dependencies

	// Global rate limiter: 100 requests per minute- use redis
	//r.Use(middlewares.RateLimiterMiddleware(100, time.Minute))

	// Group routes under /v1/2024
	public := r.Group("/v1/2024")
	{
		public.GET("/ping", pongController.PingHandler)
	}

	// Protected routes with middleware
	protected := r.Group("/v1/2024", middlewares.AuthenMiddleware())
	{
		protected.GET("/hello/:name", userController.HelloByNameHandler)
		protected.GET("/users", userController.GetUserInfoHandler)
		protected.GET("/user/:id", userController.GetUserByID)
	}

	return r
}
