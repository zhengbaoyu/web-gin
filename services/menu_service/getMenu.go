package menu_service

import (
	"strconv"
	"web-gin/common/pkg/db"
	"web-gin/models"
)

//获取功能列表
func GetInfoList() (err error, list interface{}) {
	var menuList []models.Menu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList
}

func getBaseMenuTreeMap() (err error, treeMap map[string][]models.Menu) {
	var allMenus []models.Menu
	treeMap = make(map[string][]models.Menu)
	err = db.DB.Order("sort", true).Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func getBaseChildrenList(menu *models.Menu, treeMap map[string][]models.Menu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}