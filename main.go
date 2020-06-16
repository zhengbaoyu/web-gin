package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web-gin/common/conf"
	"web-gin/common/pkg/db"
	"web-gin/routers"
)

func init() {
	//加载配置文件
	conf.GetConfig()
	//加载数据库
	db.GetDB()
	//加载redis
}

func main() {
	gin.SetMode(conf.Config.Server.RunMode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", conf.Config.Server.HttpPort)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    conf.Config.Server.ReadTimeout * time.Second,
		WriteTimeout:   conf.Config.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
