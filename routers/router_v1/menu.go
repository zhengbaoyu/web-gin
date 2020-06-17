package router_v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/controllers/api/api_v1"
	"web-gin/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.AuthJWT())
	{
		MenuRouter.GET("list", api_v1.GetMenuList)        // 权限菜单列表
		MenuRouter.POST("add", api_v1.AddMenu)            // 新增权限菜单
		MenuRouter.GET("view/:id", api_v1.ViewMenu)        // 查看权限菜单
		MenuRouter.PUT("update/:id", api_v1.UpdateMenu)    // 修改权限菜单
		MenuRouter.DELETE("delete/:id", api_v1.DeleteMenu) // 删除权限菜单
	}
	return MenuRouter
}
