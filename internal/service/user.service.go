package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetInfoUserService(c *gin.Context) {
	info := us.userRepo.GetInfoUserRepo()
	c.JSON(http.StatusOK, gin.H{"user_info": info})
}
