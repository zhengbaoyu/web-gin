package router_v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/controllers/api/api_v1"
	"web-gin/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.AuthJWT())
	{
		UserRouter.GET("/info/:id", api_v1.GetUser)
	}
}
