package initialization

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func InitRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Username: viper.GetString("redis.username"),
		Password: viper.GetString("redis.password"),
	})
	err := redisClient.Ping(context.TODO()).Err()
	if err != nil {
		panic(err.Error())
	}
	return redisClient
}
