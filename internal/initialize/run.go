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

	//InitMysql()
	InitMysqlC()

	InitRedis()

	InitKafka()
	defer CloseKafka()

	r := InitRouter()

	port := ":8081"
	if err := r.Run(port); err != nil {
		global.AppLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
