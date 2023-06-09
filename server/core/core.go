package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"server/cache"
	"server/dao"
	"server/initialization"
	"server/middleware"
	"server/router"
)

func InitCore() {
	initialization.InitViper()
	cache.InitRedisClient(initialization.InitRedis())
	dao.InitMysqlClient(initialization.InitMysql())

	app := gin.Default()
	app.Use(middleware.Cors())
	app.StaticFS("/static/image", http.Dir("./static/image"))
	group := app.Group("/api")
	router.InitRoutes(group)
	app.Run(fmt.Sprintf(viper.GetString("other.serverUrl")))
}
