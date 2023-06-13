package cache

import (
	"context"
	"fmt"
	"log"
	"server/global"
)

func tokenKey(uid int) string {
	return fmt.Sprintf("%s:%d", "token", uid)
}

func SaveToken(token string, uid int) {
	redisClient.Set(context.Background(), tokenKey(uid), token, global.REDISSAVETOKENTTL)
}

func DelToken(uid int) {
	redisClient.Del(context.Background(), tokenKey(uid))
}

func GetToken(uid int) string {
	result, err := redisClient.Get(context.Background(), tokenKey(uid)).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	return result
}

func VerifyTokenIsExist(uid int, token string) bool {
	result, err := redisClient.Get(context.Background(), tokenKey(uid)).Result()
	if err != nil || result == "" {
		log.Println(err)
		return false
	} else if result != token {
		return false
	}
	return true
}
