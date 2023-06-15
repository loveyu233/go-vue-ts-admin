package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/dao"
	"server/domain/model"
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
		return
	}
	c.JSON(200, response.NewSuccessResponse(carousels))
}

func (CarouselApi) GetShowCarousels(c *gin.Context) {
	carousels, err := dao.GetShowCarousels()
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.CarouselsGetErr))
		return
	}
	c.JSON(200, response.NewSuccessResponse(carousels))
}

func (CarouselApi) AddCarousel(c *gin.Context) {
	var ca model.Carousel
	err := c.ShouldBindJSON(&ca)
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.ParameterValidationFailed))
		return
	}
	row, err := dao.AddCarousel(ca)
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.CarouselsAddErr))
		return
	}
	c.JSON(200, response.NewSuccessResponse(row))
}

func (CarouselApi) DeleteCarousel(c *gin.Context) {
	id := c.DefaultQuery("id", "-1")
	if id == "-1" {
		c.JSON(200, response.NewErrorResponse(valid.ParameterValidationFailed))
		return
	}
	row, err := dao.DeleteCarousel(id)
	if err != nil {
		log.Println(err)
		c.JSON(200, response.NewErrorResponse(valid.CarouselsDeleteErr))
		return
	}
	c.JSON(200, response.NewSuccessResponse(row))
}

func (CarouselApi) UpdateCarousel(c *gin.Context) {
	var carousel model.Carousel
	err := c.ShouldBindJSON(&carousel)
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.ParameterValidationFailed))
		return
	}
	row, err := dao.UpdateCarousel(carousel)
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.CarouselsUpdateErr))
		return
	}
	c.JSON(200, response.NewSuccessResponse(row))
}

// OperateCarouselIn TODO 有问题
func (CarouselApi) OperateCarouselIn(c *gin.Context) {
	array, ok := c.GetQueryArray("ids[]")
	if !ok {
		c.JSON(200, response.NewErrorResponse(valid.ParameterValidationFailed))
		return
	}
	operate := c.DefaultQuery("operate", "-1")
	switch operate {
	case "1":
		for i := range array {
			dao.DeleteCarousel(array[i])
		}
	case "2":
		for i := range array {
			dao.UpdateCarouselShow(1, array[i])
		}
	case "3":
		for i := range array {
			dao.UpdateCarouselShow(0, array[i])
		}
	default:
		//错误
		c.JSON(200, response.NewErrorResponse(valid.ParameterValidationFailed))
		return
	}
	c.JSON(200, response.NewSuccessResponse(len(array)))
}
