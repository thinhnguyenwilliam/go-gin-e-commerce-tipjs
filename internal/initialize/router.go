package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"github.com/thinhcompany/ecommerce-ver-2/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.ConfigGlobal.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		// Optionally add your own middleware
		r.Use(gin.Recovery())
	}

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/v1/2024")

	// Health check route
	mainGroup.GET("/check-status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Initialize user and product routes
	userRouter.InitUserRouter(mainGroup)
	userRouter.InitProductRouter(mainGroup)

	// Initialize admin routes
	managerRouter.InitUserRouter(mainGroup)
	managerRouter.InitAdminRouter(mainGroup)

	return r
}

// func InitRouter() *gin.Engine {
// 	// Create a new Gin router
// 	r := gin.Default()

// 	// Step 1: Setup dependencies manually
// 	userRepo := repo.NewUserRepo()
// 	userService := service.NewUserService(userRepo)
// 	userController := c.NewUserController(userService)
// 	pongController := c.NewPongController() // if your PongController has no dependencies

// 	// Global rate limiter: 100 requests per minute- use redis
// 	//r.Use(middlewares.RateLimiterMiddleware(100, time.Minute))

// 	// Group routes under /v1/2024
// 	public := r.Group("/v1/2024")
// 	{
// 		public.GET("/ping", pongController.PingHandler)
// 	}

// 	// Protected routes with middleware
// 	protected := r.Group("/v1/2024", middlewares.AuthenMiddleware())
// 	{
// 		protected.GET("/hello/:name", userController.HelloByNameHandler)
// 		protected.GET("/users", userController.GetUserInfoHandler)
// 		protected.GET("/user/:id", userController.GetUserByID)
// 	}

// 	return r
// }
