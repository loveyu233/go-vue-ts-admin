package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/cache"
	"server/dao"
	"server/domain/dto"
	"server/domain/response"
	"server/domain/valid"
)

type MenuService struct {
}

func NeMenuService() *MenuService {
	return new(MenuService)
}

func (m MenuService) GetMenus() *response.APIResponse {
	menus, err := dao.GetMenusInfo()
	if err != nil {
		log.Println(err)
		return response.NewErrorResponse(valid.ServerErrorCode)
	}
	return response.NewSuccessResponse(menus)
}

func (m MenuService) GetMenusLimit(page, pageSize int) *response.APIResponse {
	menus, total := dao.GetMenusLimit(page, pageSize)
	return response.NewSuccessResponse(gin.H{
		"menus": menus,
		"total": total,
	})
}

func (m MenuService) MenuUpdate(menu dto.MenuUpdate) *response.APIResponse {
	affected := dao.MenuUpdate(menu)
	if affected <= 0 {
		return response.NewErrorResponse(valid.MenuInfoUpdateErr)
	}
	cache.DelMenu()
	return response.NewSuccessResponse(gin.H{
		"msg": "ok",
	})
}
