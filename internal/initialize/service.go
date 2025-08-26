// internal/initialize/service.go
package initialize

import (
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"github.com/thinhcompany/ecommerce-ver-2/internal/db"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service/impl"
)

func InitServiceInterface() {
	queries := db.New(global.Mdbc)

	// Correct initialization based on interfaces
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	// service.InitUserInfo(impl.NewUserInfoImpl(queries))     // if exists
	// service.InitUserAdmin(impl.NewUserAdminImpl(queries))   // if exists
}
