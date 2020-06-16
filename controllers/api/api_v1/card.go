package api_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-gin/common/conf"
	"web-gin/common/pkg"
	"web-gin/services/card_service"
)

func Download(c *gin.Context) {
	response := pkg.Gin{C: c}
	cardTypeId, _ := strconv.Atoi(c.Query("card_type_id"))
	cardDownloadService := card_service.CardDownload{CardTypeId: cardTypeId}
	fileName, err := cardDownloadService.CardDownload()
	if err != nil {
		response.ResponseOk(http.StatusOK, pkg.ERR_DOWNLOAD_FAIL, err)
	}
	ExcelFullUrl := conf.Config.App.PrefixUrl + "/" + conf.Config.Export.ExportSavePath + fileName
	GetExcelPath := conf.Config.Export.ExportSavePath + fileName
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, gin.H{
		"export_url":      ExcelFullUrl,
		"export_save_url": GetExcelPath,
	})
}
