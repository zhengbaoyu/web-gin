package role_repository

import (
	"web-gin/common/pkg/db"
	"web-gin/controllers/api/api_v1/check_form"
	"web-gin/models"
)

//获取角色列表
func GetList(pageInfo check_form.PageInfo, wheres map[string]interface{}) (err error, list []models.Roles, total int) {
	err = db.DB.Limit(pageInfo.Page).Offset(pageInfo.PageSize).Where(wheres).Find(&list).Error
	db.DB.Model(&models.Roles{}).Where("parent_id = 0").Count(&total)
	return err, list, total
}

//获取子角色
func GetChildrenList(id uint, authority *models.Roles) error {
	err := db.DB.Where("parent_id = ? and delete_status = 0", id).Find(&authority.Children).Error
	if err != nil {
		return err
	}
	return nil
}

//添加
func Add(roles models.Roles) error {
	err := db.DB.Create(roles).Error
	if err != nil {
		return err
	}
	return nil
}
