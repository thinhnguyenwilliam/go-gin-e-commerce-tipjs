package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/routers"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Register routes from routers package
	routers.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
