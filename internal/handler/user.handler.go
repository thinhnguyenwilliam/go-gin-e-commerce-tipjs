package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
	"github.com/thinhcompany/ecommerce-ver-2/internal/vo"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
)

type UserHandler struct {
	userService service.IUserService
}

// Constructor
func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GET /user/check-email?email=hoang@gmail.com
func (h *UserHandler) CheckEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(response.ErrorCodeParamInvalid, nil))
		return
	}
	resp := h.userService.CheckUserExists(email)
	c.JSON(http.StatusOK, resp)
}

// Register user
func (h *UserHandler) Register(c *gin.Context) {
	var req vo.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": err.Error(),
		// })
		c.JSON(http.StatusBadRequest, response.ErrorResponse(response.ErrorCodeParamInvalid, nil))
		return
	}

	// Log the request body
	log.Printf("Received registration request: email=%s, purpose=%s", req.Email, req.Purpose)

	jsonData, _ := json.Marshal(req)
	log.Println("Request body:", string(jsonData))

	resp := h.userService.Register(req.Email, req.Purpose)
	c.JSON(http.StatusOK, resp)
}
