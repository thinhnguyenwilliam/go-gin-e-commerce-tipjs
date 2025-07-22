package repo

import (
	"log"

	"github.com/thinhcompany/ecommerce-ver-2/global"
	"github.com/thinhcompany/ecommerce-ver-2/internal/db"
)

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct {
	sqlc *db.Queries
}

// Constructor to create a new UserRepo instance
func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: db.New(global.Mdbc),
	}
}

func (r *userRepo) GetUserByEmail(email string) bool {
	user, err := r.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		log.Printf("User not found or SQL error: %v", err)
		return false
	}
	log.Printf("User found: %s with ID: %d", user.UsrEmail, user.UsrID)
	return true
}
