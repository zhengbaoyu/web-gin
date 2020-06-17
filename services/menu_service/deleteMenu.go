package menu_service

import (
	"web-gin/common/pkg"
	"web-gin/common/pkg/db"
	"web-gin/models"
	"web-gin/repositorys/menu_repository"
)

//删除权限功能
func DeleteMenu(id uint) int {
	//校验
	var menuData models.Menus
	db.DB.Where("parent_id=? and delete_status=?", id, 0).First(&menuData)
	if menuData.ID > 0 {
		return pkg.ERR_DELETE_MENU_EXIST
	}
	delErr := menu_repository.Delete(id)
	if delErr != nil {
		return pkg.ERR_DELETE_MENU_FAIL
	}
	return pkg.SUCCESS
}
