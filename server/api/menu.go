package api

import (
	"github.com/gin-gonic/gin"
	"server/cache"
	"server/domain/dto"
	"server/domain/response"
	"server/domain/valid"
	"server/service"
	"strconv"
)

type MenuApi struct {
}

func NewMenuApi() *MenuApi {
	return new(MenuApi)
}

var menuService = service.NeMenuService()

func (m MenuApi) GetMenusTree(c *gin.Context) {
	c.JSON(200, menuService.GetMenus())
}

func (m MenuApi) GetMenusList(c *gin.Context) {
	c.JSON(200, cache.GetMenu())
}

func (m MenuApi) MenusLimit(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "5")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	apiResponse := menuService.GetMenusLimit(page, pageSize)
	c.JSON(200, apiResponse)
}

func (m MenuApi) MenusUpdate(c *gin.Context) {
	var menuDto dto.MenuUpdate
	err := c.ShouldBindJSON(&menuDto)
	if err != nil {
		c.JSON(200, response.NewErrorResponse(valid.MenuInfoUpdateErr))
		return
	}
	c.JSON(200, menuService.MenuUpdate(menuDto))
}
