package card_service

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
	"web-gin/common/conf"
	"web-gin/common/pkg/file"
	"web-gin/repositorys/card_repository"
)

type CardDownload struct {
	CardTypeId int
}

func (down *CardDownload) CardDownload() (string, error) {
	fmt.Println("CardTypeId:", down.CardTypeId)
	datas, err := card_repository.GetDownloadCard(down.CardTypeId)
	if err != nil {
		return "", err
	}
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("标签信息")
	if err != nil {
		return "", err
	}
	titles := []string{"卡号", "密码", "二维码"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}
	for _, v := range datas {
		values := []string{
			v.CardIdentifyNum,
			v.CardPwd,
			v.QrCode,
		}
		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}
	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "card-" + time + conf.Config.Export.ExportFileNameSuffix
	dirFullPath := conf.Config.App.RuntimeRootPath + conf.Config.Export.ExportSavePath
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}
	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}
