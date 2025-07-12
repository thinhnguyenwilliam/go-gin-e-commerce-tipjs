package initialize

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thinhcompany/ecommerce-ver-2/global"
	"go.uber.org/zap"
)

func InitRedis() {
	cfg := global.ConfigGlobal.Redis

	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	// Ping test
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := global.Rdb.Ping(ctx).Result()
	if err != nil {
		global.AppLogger.Fatal("Redis connection failed", zap.Error(err))
	} else {
		global.AppLogger.Info("Connected to Redis",
			zap.String("addr", cfg.Addr),
			zap.Int("db", cfg.DB),
		)
	}
}
