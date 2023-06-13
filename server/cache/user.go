package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"server/domain/vo"
)

type UserCache struct {
}

func SaveUserInfo(uid int, value vo.UserVo) {
	userJson, _ := json.Marshal(&value)
	set(fmt.Sprintf("%s%d", viper.GetString("other.redisuserprefix"), uid), userJson)
}

func GetUserInfo(uid int) *vo.UserVo {
	var userVo vo.UserVo
	userJson := get(fmt.Sprintf("%s%d", viper.GetString("other.redisuserprefix"), uid))
	if userJson == "" {
		return nil
	}
	json.Unmarshal([]byte(userJson), &userVo)
	return &userVo
}

func DeleteUserInfoData(uid int) {
	redisClient.Del(context.Background(), fmt.Sprintf("%s%d", viper.GetString("other.redisuserprefix"), uid))
	DelToken(uid)
}
