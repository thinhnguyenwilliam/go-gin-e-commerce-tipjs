package redisutil

import (
	"context"
	"fmt"
	"time"

	"github.com/thinhcompany/ecommerce-ver-2/global"
)

func IsOtpStillValid(hashedEmail, purpose string) (bool, error) {
	otpKey := fmt.Sprintf("usr:%s:otp", hashedEmail)
	exists, err := global.Rdb.Exists(context.Background(), otpKey).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// SetKey sets a key with a value and TTL
func SetKey(ctx context.Context, key string, value string, ttl time.Duration) error {
	return global.Rdb.Set(ctx, key, value, ttl).Err()
}

// GetKey gets the value of a key
func GetKey(ctx context.Context, key string) (string, error) {
	return global.Rdb.Get(ctx, key).Result()
}

func DelKey(ctx context.Context, key string) error {
	return global.Rdb.Del(ctx, key).Err()
}

func ExistsKey(ctx context.Context, key string) (bool, error) {
	result, err := global.Rdb.Exists(ctx, key).Result()
	return result > 0, err
}

// func Example() {
// 	ctx := context.Background()
// 	err := redisutil.SetKey(ctx, "myKey", "some value", 5*time.Minute)
// 	if err != nil {
// 		fmt.Println("Set failed:", err)
// 	}

// 	val, err := redisutil.GetKey(ctx, "myKey")
// 	if err != nil {
// 		fmt.Println("Get failed:", err)
// 	} else {
// 		fmt.Println("Value:", val)
// 	}
// }
