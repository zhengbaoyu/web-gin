package conf

import (
	"fmt"
	"github.com/jinzhu/configor"
	"time"
)

var Config = struct {
	APPName string
	//应用配置
	App struct {
		PrefixUrl       string
		RuntimeRootPath string
	}
	//服务器配置
	Server struct {
		RunMode      string
		HttpPort     int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	//数据库
	Database struct {
		Type        string
		User        string
		Password    string
		Host        string
		Name        string
		TablePrefix string
	}
	//日志
	Log struct {
		LogSavePath string
		LogSaveName string
		LogFileExt  string
		TimeFormat  string
	}
	//RabbitMq 队列
	RabbitMq struct {
		User     string
		Password string
		Host     string
	}
	//数据导出
	Export struct {
		ExportSavePath       string
		ExportFileNameSuffix string
	}
}{}

func GetRabbitMqConf() string {
	return fmt.Sprintf("amqp://%s:%s@%s",
		Config.RabbitMq.User,
		Config.RabbitMq.Password,
		Config.RabbitMq.Host)
}

func GetConfig() {
	configor.Load(&Config, "common/conf/config.yml")
}
