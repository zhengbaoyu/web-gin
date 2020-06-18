package role_service

import (
	"web-gin/common/pkg"
	"web-gin/common/pkg/db"
	"web-gin/models"
	"web-gin/repositorys/role_repository"
)

type RoleAddService struct {
	ParentId      int
	AuthorityName string
}

func (r *RoleAddService) AddRole() int {
	//校验
	var roleData models.Roles
	db.DB.Where("authority_name = ?", r.AuthorityName).First(&roleData)
	if roleData.ID > 0 {
		return pkg.ERR_ROLE_NAME_EXIST
	}
	roleM := models.Roles{
		AuthorityName: r.AuthorityName,
		ParentId:      r.ParentId,
	}
	err := role_repository.Add(roleM)
	if err != nil {
		return pkg.ERR_ADD_ROLE_FAIL
	}
	return pkg.SUCCESS
}
