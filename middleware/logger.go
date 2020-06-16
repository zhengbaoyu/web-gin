package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
	"web-gin/common/pkg/rabbitmq"
)

func LoggerToFile() gin.HandlerFunc {
	//实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		msg := logger.WithFields(logrus.Fields{
			"statusCode":  c.Writer.Status(),
			"latencyTime": latencyTime,
			"clientIp":    c.ClientIP(),
			"method":      c.Request.Method,
			"requestUrl":  c.Request.RequestURI,
			"params":      c.Request.URL.RawQuery,
		})
		LogMQ(msg.Data)
		//fmt.Println("log详情:", msg.Data)
	}
}

//队列
func LogMQ(data map[string]interface{}) {
	logjson, _ := json.Marshal(data)
	mString := string(logjson)
	Publish := rabbitmq.NewRabbitMQSimple("logs")
	Publish.PublishSimple(mString)
}
