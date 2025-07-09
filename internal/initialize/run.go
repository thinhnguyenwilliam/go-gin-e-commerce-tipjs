package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8080") // Note: capital R in Run()
}
