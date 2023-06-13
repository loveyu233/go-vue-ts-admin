package router

import "github.com/gin-gonic/gin"

func InitRoutes(group *gin.RouterGroup) {
	InitUserRoute(group)
	InitMenuRouter(group)
	InitCarouselRouter(group)
}
