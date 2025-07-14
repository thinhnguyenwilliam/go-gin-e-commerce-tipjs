package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct definition
type PongController struct{}

// Constructor function
func NewPongController() *PongController {
	return &PongController{}
}

// Handler method
func (p *PongController) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
