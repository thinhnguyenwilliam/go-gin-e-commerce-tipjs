// internal/service/impl/user_login_impl.go
package impl

import (
	"context"

	"github.com/thinhcompany/ecommerce-ver-2/internal/db"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
)

type sUserLogin struct {
	r *db.Queries
}

// Constructor
func NewUserLoginImpl(r *db.Queries) service.IUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// Implements IUserLogin interface
func (s *sUserLogin) Login(ctx context.Context) error {
	// TODO: Add login logic here
	return nil
}
