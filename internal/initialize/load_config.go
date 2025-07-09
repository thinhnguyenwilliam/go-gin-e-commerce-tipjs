package initialize

import (
	"log"

	"github.com/spf13/viper"
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/setting"
)

func LoadConfig() {
	v := viper.New()
	v.SetConfigName("local")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var cfg setting.Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	global.ConfigGlobal = &cfg

	// Debug output
	port := global.ConfigGlobal.Server.Port
	dbUser := global.ConfigGlobal.MySQL.User
	log.Printf("Config loaded: server port = %d, db user = %s", port, dbUser)
}
