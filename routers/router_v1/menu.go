package router_v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/controllers/api/api_v1"
	"web-gin/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.AuthJWT())
	{
		MenuRouter.GET("getMenuList", api_v1.GetMenuList)           // 分页获取基础menu列表
	}
	return MenuRouter
}
