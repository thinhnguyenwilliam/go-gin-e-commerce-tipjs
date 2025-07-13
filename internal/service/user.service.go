package service

import (
	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) response.ResponseData
}

type userService struct {
	userRepo repo.IUserRepo // ✅ just the interface, no pointer
}

// Constructor for userService
func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Example implementation
func (s *userService) Register(email string, purpose string) response.ResponseData {
	if s.userRepo.GetUserByEmail(email) {
		return response.ErrorResponse(response.ErrorCodeUserHasExists, nil)
	}

	// Continue registration logic here...

	return response.SuccessResponse("Register success")
}

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService(userRepo *repo.UserRepo) *UserService {
// 	return &UserService{userRepo: userRepo}
// }

// func (us *UserService) GetInfoUserService(c *gin.Context) {
// 	info := us.userRepo.GetInfoUserRepo()
// 	c.JSON(http.StatusOK, gin.H{"user_info": info})
// }
