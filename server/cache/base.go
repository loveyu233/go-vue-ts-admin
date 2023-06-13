package cache

import "github.com/go-redis/redis/v8"

var redisClient *redis.Client

func InitRedisClient(client *redis.Client) {
	redisClient = client
}
