package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/thinhcompany/ecommerce-ver-2/global"
	"go.uber.org/zap"
)

func InitMysqlC() error {
	mysqlCfg := global.ConfigGlobal.MySQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		global.AppLogger.Error("Database connection failed",
			zap.String("host", mysqlCfg.Host),
			zap.Error(err),
		)
		return err
	}

	// Set SQL pool
	db.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	db.SetMaxOpenConns(mysqlCfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(mysqlCfg.ConnMaxLifetime) * time.Second)

	global.Mdbc = db // ✅ only use Mdbc for SQLC

	global.AppLogger.Info("Init MySQL success",
		zap.String("host", mysqlCfg.Host),
		zap.Int("port", mysqlCfg.Port),
		zap.String("user", mysqlCfg.User),
		zap.String("db", mysqlCfg.DBName),
		zap.Int("maxIdleConns", mysqlCfg.MaxIdleConns),
		zap.Int("maxOpenConns", mysqlCfg.MaxOpenConns),
		zap.Int("connMaxLifetime(sec)", mysqlCfg.ConnMaxLifetime),
	)

	return nil
}
