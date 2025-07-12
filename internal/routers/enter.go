package routers

import (
	"github.com/thinhcompany/ecommerce-ver-2/internal/routers/manager"
	"github.com/thinhcompany/ecommerce-ver-2/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

// Global instance of all router groups
var RouterGroupApp = new(RouterGroup)
