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
		c.JSON(200, gin.H{"status": "ok honey"})
	})

	// Initialize user and product routes
	userRouter.InitUserRouter(mainGroup)
	userRouter.InitProductRouter(mainGroup)

	// Initialize admin routes
	managerRouter.InitUserRouter(mainGroup)
	managerRouter.InitAdminRouter(mainGroup)

	return r
}
