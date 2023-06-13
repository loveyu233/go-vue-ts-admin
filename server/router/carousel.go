package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func InitCarouselRouter(group *gin.RouterGroup) {
	carousel := group.Group("/carousel")
	carouselApi := api.NewCarouselApi()
	carousel.GET("/all", carouselApi.GetCarousels)
}
