package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func InitMenuRouter(group *gin.RouterGroup) {
	menuApi := api.NewMenuApi()
	menu := group.Group("/menu")
	// group.Use(middleware.UserAuth())
	menu.GET("/info", menuApi.GetMenusTree)
	menu.GET("/list", menuApi.GetMenusList)
	menu.GET("/limit", menuApi.MenusLimit)
	menu.POST("/update", menuApi.MenusUpdate)
}
