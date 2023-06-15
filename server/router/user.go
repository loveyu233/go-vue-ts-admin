package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

var userApi = api.NewUserApi()

func InitUserRoute(group *gin.RouterGroup) {
	user := group.Group("/user")
	user.POST("/login", userApi.VerifyUsernamePassword)

	auth := user.Group("")
	auth.Use(middleware.UserAuth())
	auth.GET("/info", userApi.GetUserInfo)
	auth.GET("/verify", userApi.VerifyToken)
	auth.GET("/exit", userApi.UserExit)
	auth.POST("/image", userApi.AddUserIcon)
	auth.POST("/update", userApi.UpdateUserInfo)
}
