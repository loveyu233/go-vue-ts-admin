package cache

import (
	"context"
	"encoding/json"
	"log"
	"server/domain/model"
)

func SaveMenu(menus []*model.SysBaseMenus) {
	marshal, _ := json.Marshal(menus)
	redisClient.Set(context.Background(), "menus", string(marshal), -1)
}

func GetMenu() []model.SysBaseMenus {
	result, err := redisClient.Get(context.Background(), "menus").Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	var menus = make([]model.SysBaseMenus, 0)
	err = json.Unmarshal([]byte(result), &menus)
	if err != nil {
		log.Println(err)
		return nil
	}
	return menus
}

func DelMenu() {
	redisClient.Del(context.Background(), "menu")
}
