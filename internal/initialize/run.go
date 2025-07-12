package initialize

import (
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()

	InitLogger()
	defer global.AppLogger.Sync() // flush logs on exit
	global.AppLogger.Info("Config log ok!!", zap.String("ok", "success"))

	InitMysql()

	InitRedis()

	r := InitRouter()

	if err := r.Run(":8080"); err != nil {
		global.AppLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
