package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-gin/common/conf"
	"web-gin/controllers/api/api_v1"
	"web-gin/middleware"
	"web-gin/routers/router_v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerToFile())
	//下载Excel
	r.StaticFS("/export", http.Dir(conf.Config.App.RuntimeRootPath+conf.Config.Export.ExportSavePath))

	v1 := r.Group("/v1")
	//下载会员卡数据
	v1.GET("/card/download", api_v1.Download)
	//注册
	v1.POST("/user/register", api_v1.Register)
	//登录
	v1.POST("/user/login", api_v1.Login)
	//用户
	router_v1.InitUserRouter(v1)
	//权限菜单
	router_v1.InitMenuRouter(v1)
	return r
}
