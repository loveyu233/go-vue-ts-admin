package initialization

import (
	"github.com/spf13/viper"
)

func InitViper() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
		return
	}
	err = viper.WriteConfig()
	if err != nil {
		panic(err.Error())
		return
	}
}
