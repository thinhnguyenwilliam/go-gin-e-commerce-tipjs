package test

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9" // ✅ import this!
	"github.com/stretchr/testify/assert"
	"github.com/thinhcompany/ecommerce-ver-2/global"
)

// func init() {
// 	initialize.LoadConfig()   // If needed
// 	initialize.InitLogger()   // Optional
// 	initialize.InitRedis()    // ✅ Initialize Redis
// }

func init() {
	// Only initialize Redis if not already done
	if global.Rdb == nil {
		global.Rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // or load from config
			Password: "",               // no password set
			DB:       0,
		})

		// Test connection
		ctx := context.Background()
		if err := global.Rdb.Ping(ctx).Err(); err != nil {
			panic("Redis not connected: " + err.Error())
		}
	}
}

func TestRedisSetGet(t *testing.T) {
	ctx := context.Background()
	key := "test:anhiu"
	value := "hello e iu"

	// SET key with TTL
	err := global.Rdb.Set(ctx, key, value, 10*time.Second).Err()
	assert.NoError(t, err, "Redis SET failed")

	// GET key
	result, err := global.Rdb.Get(ctx, key).Result()
	assert.NoError(t, err, "Redis GET failed")
	assert.Equal(t, value, result, "GET value mismatch")

	// Cleanup
	global.Rdb.Del(ctx, key)

	t.Run("test expire", func(t *testing.T) {
		err := global.Rdb.Set(ctx, "expire:key", "123", 2*time.Second).Err()
		assert.NoError(t, err)

		time.Sleep(3 * time.Second)

		val, err := global.Rdb.Get(ctx, "expire:key").Result()
		assert.ErrorIs(t, err, redis.Nil) // ✅ Correct nil comparison
		assert.Empty(t, val)
	})
}
