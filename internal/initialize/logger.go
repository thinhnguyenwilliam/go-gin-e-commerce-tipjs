package initialize

import (
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/logger"
)

func InitLogger() {
	global.AppLogger = logger.NewLogger(&global.ConfigGlobal.Logger)
}
