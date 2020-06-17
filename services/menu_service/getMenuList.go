package menu_service

import (
	"strconv"
	"web-gin/common/pkg/db"
	"web-gin/models"
)

//获取功能列表
func GetMenuList() (err error, list interface{}) {
	var menuList []models.Menus
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList
}

func getBaseMenuTreeMap() (err error, treeMap map[string][]models.Menus) {
	var allMenus []models.Menus
	treeMap = make(map[string][]models.Menus)
	err = db.DB.Order("sort", true).Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func getBaseChildrenList(menu *models.Menus, treeMap map[string][]models.Menus) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}