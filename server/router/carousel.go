package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func InitCarouselRouter(group *gin.RouterGroup) {
	carousel := group.Group("/carousel")
	carouselApi := api.NewCarouselApi()
	carousel.GET("/all", carouselApi.GetCarousels)
	carousel.GET("/show", carouselApi.GetShowCarousels)
	carousel.POST("/add", carouselApi.AddCarousel)
	carousel.GET("/del", carouselApi.DeleteCarousel)
	carousel.POST("/update", carouselApi.UpdateCarousel)
	carousel.GET("/delete", carouselApi.OperateCarouselIn)
}
