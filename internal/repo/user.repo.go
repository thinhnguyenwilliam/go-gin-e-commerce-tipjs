package repo

import (
	"log"

	"github.com/thinhcompany/ecommerce-ver-2/global"
	"github.com/thinhcompany/ecommerce-ver-2/internal/model"
)

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

// Constructor to create a new UserRepo instance
func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (r *userRepo) GetUserByEmail(email string) bool {
	var user model.GoCrmUser
	result := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&user)
	log.Printf("Checked email %s - Exists: %v", email, result.RowsAffected != NumberNil)
	return result.RowsAffected != NumberNil
}
