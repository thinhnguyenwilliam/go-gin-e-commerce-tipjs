//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/thinhcompany/ecommerce-ver-2/internal/handler"
	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
)

func InitUserHandler() *handler.UserHandler {
	wire.Build(
		repo.NewUserAuthRepository,
		repo.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,
	)
	return &handler.UserHandler{}
}
