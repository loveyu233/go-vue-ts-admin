package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"mime/multipart"
	"server/cache"
	"server/dao"
	"server/domain/dto"
	"server/domain/response"
	"server/domain/valid"
	"server/domain/vo"
	"server/jwt"
	"server/utils/update"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

type M map[string]interface{}

func (u UserService) VerifyUsernamePassword(userDto dto.UserDto) *response.APIResponse {
	userModel, err := dao.VerifyUsernamePassword(userDto.Username, userDto.Password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return response.NewErrorResponse(valid.UserNotFound)
		}
		log.Println(err)
		return response.NewErrorResponse(valid.ServerErrorCode)
	}

	token := cache.GetToken(userModel.ID)

	if token != "" {
		return response.NewSuccessResponse(M{"token": token})
	}
	token, err = jwt.CreateToken(userModel.ID)
	if err != nil {
		log.Println(err)
		return response.NewErrorResponse(valid.ServerErrorCode)
	}
	cache.SaveToken(token, userModel.ID)
	log.Println("createToken: ", token)
	return response.NewSuccessResponse(M{"token": token})
}

func (u UserService) GetUserInfo(uid int) *response.APIResponse {
	info := cache.GetUserInfo(uid)
	if info != nil {
		return response.NewSuccessResponse(info)
	}
	userModel, err := dao.SelectById(uid)
	if err != nil {
		return response.NewErrorResponse(valid.ServerErrorCode)
	}
	roles, err := dao.SelectRoles(uid)
	if err != nil {
		return response.NewErrorResponse(valid.ServerErrorCode)
	}
	buttons, err := dao.SelectButtons(uid)
	if err != nil {
		return response.NewErrorResponse(valid.ServerErrorCode)
	}
	routes, err := dao.SelectRoutes(uid)
	if err != nil {
		return response.NewErrorResponse(valid.ServerErrorCode)
	}
	var userVo vo.UserVo
	copier.Copy(&userVo, userModel)
	userVo.Roles = roles
	userVo.Buttons = buttons
	userVo.Routers = routes
	cache.SaveUserInfo(uid, userVo)
	return response.NewSuccessResponse(userVo)
}

func (u UserService) UserIconUpdate(file *multipart.FileHeader) *response.APIResponse {
	url := update.FileUpdate(file)
	return response.NewSuccessResponse(gin.H{
		"url": url,
	})
}

func (u UserService) UpdateUserInfo(user dto.UserUpdateDto) *response.APIResponse {
	info, err := dao.UpdateUserInfo(user)
	if err != nil || info == 0 {
		return response.NewErrorResponse(valid.UserInfoUpdateErr)
	}
	cache.DeleteUserInfoData(user.Uid)
	return response.NewSuccessResponse(gin.H{
		"msg": "修改成功",
	})
}
