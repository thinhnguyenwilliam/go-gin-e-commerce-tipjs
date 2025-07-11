package initialize

import (
	"fmt"
	"time"

	"github.com/thinhcompany/ecommerce-ver-2/global"
	model "github.com/thinhcompany/ecommerce-ver-2/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() error {
	mysqlCfg := global.ConfigGlobal.MySQL

	// DSN: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info), // change to Error or Silent in production
	})
	if err != nil {
		global.AppLogger.Error("Database connection failed",
			zap.String("host", global.ConfigGlobal.MySQL.Host),
			zap.Error(err),
		)
		return err
	}

	global.AppLogger.Info("Init MySQL success",
		zap.String("host", global.ConfigGlobal.MySQL.Host),
		zap.Int("port", global.ConfigGlobal.MySQL.Port),
		zap.String("user", global.ConfigGlobal.MySQL.User),
		zap.String("db", global.ConfigGlobal.MySQL.DBName),
		zap.Int("maxIdleConns", global.ConfigGlobal.MySQL.MaxIdleConns),
		zap.Int("maxOpenConns", global.ConfigGlobal.MySQL.MaxOpenConns),
		zap.Int("connMaxLifetime(sec)", global.ConfigGlobal.MySQL.ConnMaxLifetime),
	)
	global.Mdb = db

	SetPool()
	migratalbes()

	return nil
}

func SetPool() {
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		global.AppLogger.Fatal("Failed to get database connection pool", zap.Error(err))
	}

	mysqlCfg := global.ConfigGlobal.MySQL

	sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlCfg.ConnMaxLifetime) * time.Second)
}

func migratalbes() {
	err := global.Mdb.AutoMigrate(
		&model.User{}, // replace with your actual models
		&model.Product{},
		&model.Role{},
	)
	if err != nil {
		global.AppLogger.Fatal("Failed to auto-migrate tables", zap.Error(err))
	}
}
