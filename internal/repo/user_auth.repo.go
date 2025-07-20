package repo

import (
	"fmt"
	"time"

	"github.com/thinhcompany/ecommerce-ver-2/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp string, expiration time.Duration) error
}

type userAuthRepository struct{}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

func (u *userAuthRepository) AddOTP(email string, otp string, expiration time.Duration) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.Set(ctx, key, otp, expiration).Err()
}
