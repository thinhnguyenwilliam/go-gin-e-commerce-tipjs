package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"databases"`

	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

var AppConfig Config

func main() {
	loadConfig()

	router := gin.Default()

	router.GET("/ping/testviper", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"jwtKey":  AppConfig.Security.JWT.Key,
		})
	})

	addr := fmt.Sprintf(":%d", AppConfig.Server.Port)
	jwtKey := AppConfig.Security.JWT.Key

	log.Printf("Starting server on honey %s with JWT key: %s", addr, jwtKey)
	router.Run(addr)

}

func loadConfig() {
	viper.SetConfigName("local")    // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}

	log.Println("Config loaded successfully")
}
