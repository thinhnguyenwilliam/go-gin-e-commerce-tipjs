package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
)

// Define the controller struct
type UserController struct {
	userService *service.UserService
}

// Constructor function to create a new UserController
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{"user": "William"}))

	// or in case of error
	//c.JSON(http.StatusBadRequest, response.ErrorResponse(response.ErrorCodeParamInvalid, nil))
}

func (uc *UserController) GetUserInfoHandler(c *gin.Context) {
	uc.userService.GetInfoUserService(c)
}

// Handler method on the UserController struct
func (uc *UserController) HelloByNameHandler(c *gin.Context) {
	name := c.Param("name")              // e.g. /hello/thinh
	uid := c.DefaultQuery("uid", "0000") // ?uid=123 or default "0000"

	log.Println("Name:", name, "UID:", uid)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "thinh"},
	})
}
