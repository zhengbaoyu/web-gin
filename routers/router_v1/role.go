package router_v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/controllers/api/api_v1"
	"web-gin/middleware"
)

func InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("role").Use(middleware.AuthJWT())
	{
		MenuRouter.GET("list", api_v1.GetRoleList)         // 角色列表
		MenuRouter.POST("add", api_v1.AddRole)             // 新增角色
		MenuRouter.GET("view/:id", api_v1.ViewRole)        // 查看角色
		MenuRouter.PUT("update/:id", api_v1.UpdateRole)    // 修改角色
		MenuRouter.DELETE("delete/:id", api_v1.DeleteRole) // 删除角色
	}
	return MenuRouter
}
