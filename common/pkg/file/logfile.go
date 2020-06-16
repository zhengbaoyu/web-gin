package file

import (
	"fmt"
	"time"
	"web-gin/common/conf"
)

//获取日志文件保存目录
func GetLogFilePath() string {
	return fmt.Sprintf("%s%s", conf.Config.App.RuntimeRootPath, conf.Config.Log.LogSavePath)
}

func GetLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		conf.Config.Log.LogSaveName,
		time.Now().Format(conf.Config.Log.TimeFormat),
		conf.Config.Log.LogFileExt,
	)
}
