package router_v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/controllers/api/api_v1"
	"web-gin/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.AuthJWT())
	{
		MenuRouter.GET("menuList", api_v1.GetMenuList) // 分页获取基础menu列表
		MenuRouter.POST("addMenu", api_v1.AddMenu)     // 新增菜单
	}
	return MenuRouter
}
