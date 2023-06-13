package api

import (
	"github.com/gin-gonic/gin"
	"server/dao"
	"server/domain/response"
	"server/domain/valid"
)

type CarouselApi struct {
}

func NewCarouselApi() *CarouselApi {
	return new(CarouselApi)
}

func (CarouselApi) GetCarousels(c *gin.Context) {
	carousels, err := dao.GetCarousels()
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.CarouselsGetErr))
	}
	c.JSON(200, response.NewSuccessResponse(carousels))
}
