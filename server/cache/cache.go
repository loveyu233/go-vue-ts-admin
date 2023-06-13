package cache

import (
	"context"
	. "server/global"
)

func get(key string) string {
	result, _ := redisClient.Get(context.Background(), key).Result()
	return result
}

func set(key string, value interface{}) {
	redisClient.Set(context.Background(), key, value, REDISTTL)
}
