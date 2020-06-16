package pkg

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

//定义返回数据结构

type Response struct {
	Status     bool        `json:"status"`
	ReturnCode int         `json:"returnCode"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func (g *Gin) ResponseOk(httpCode, ReturnCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Status:     true,
		ReturnCode: ReturnCode,
		Msg:        GetMsg(ReturnCode),
		Data:       data,
	})
}

func (g *Gin) ResponseErr(httpCode, ReturnCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Status:     false,
		ReturnCode: ReturnCode,
		Msg:        GetMsg(ReturnCode),
		Data:       data,
	})
}
