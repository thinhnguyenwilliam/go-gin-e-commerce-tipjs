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

	// if err := InitMysql(); err != nil {
	// 	global.AppLogger.Fatal("Failed to initialize MySQL", zap.Error(err))
	// }

	// if err := InitRedis(); err != nil {
	// 	global.AppLogger.Fatal("Failed to initialize Redis", zap.Error(err))
	// }

	r := InitRouter()

	if err := r.Run(":8080"); err != nil {
		global.AppLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
