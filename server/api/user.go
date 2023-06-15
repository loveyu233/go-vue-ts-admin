package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"server/cache"
	"server/domain/dto"
	"server/domain/response"
	"server/domain/valid"
	"server/service"
)

type UserApi struct {
}

func NewUserApi() *UserApi {
	return new(UserApi)
}

var userService = service.NewUserService()

func (u UserApi) VerifyUsernamePassword(c *gin.Context) {
	var user dto.UserDto
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.ParameterValidationFailed))
		return
	}
	c.JSON(200, userService.VerifyUsernamePassword(user))
}

func (u UserApi) GetUserInfo(c *gin.Context) {
	uid := c.GetInt("uid")
	info := userService.GetUserInfo(uid)
	c.JSON(200, info)
}

func (u UserApi) VerifyToken(c *gin.Context) {
	c.JSON(200, response.NewSuccessResponse(gin.H{"verify": true}))
}

func (u UserApi) UserExit(c *gin.Context) {
	uid := c.GetInt("uid")
	cache.DelToken(uid)
	c.JSON(200, response.NewSuccessResponse(gin.H{"message": "退出成功"}))
}

const dst = "./static/image/"

func (u UserApi) AddUserIcon(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.IconUpdateErr))
		return
	}
	err = c.SaveUploadedFile(file, fmt.Sprintf("%s%s", dst, file.Filename))
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.IconUpdateErr))
		return
	}
	var serverUrl = fmt.Sprintf("%s%s", "http://", viper.GetString("other.serverUrl"))

	//userService.UserIconUpdate(file)
	c.JSON(200, response.NewSuccessResponse(gin.H{
		"url": fmt.Sprintf("%s%s%s", serverUrl, "/static/image/", file.Filename),
	}))
}

func (u UserApi) UpdateUserInfo(c *gin.Context) {
	uid := c.GetInt("uid")
	var userDto dto.UserUpdateDto
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		log.Println(err)
		c.JSON(200, response.NewErrorResponse(valid.UserInfoUpdateErr))
		return
	}
	userDto.Uid = uid
	apiResponse := userService.UpdateUserInfo(userDto)
	c.JSON(200, apiResponse)
}
