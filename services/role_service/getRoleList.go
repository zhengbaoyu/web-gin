package role_service

import (
	"web-gin/controllers/api/api_v1/check_form"
	"web-gin/models"
	"web-gin/repositorys/role_repository"
)

type RoleListService struct {
	Page          int
	PageSize      int
	AuthorityName string
}

//获取角色列表
func (r *RoleListService) GetRoleList() (error, interface{}, int) {
	pageInfo := check_form.PageInfo{
		Page:     r.PageSize,
		PageSize: r.PageSize * (r.Page - 1),
	}
	err, list, total := role_repository.GetList(pageInfo, r.GetWheres())
	if len(list) > 0 {
		for k := range list {
			err = findChildrenAuthority(&list[k])
		}
	}
	return err, list, total
}

//查看子角色
func findChildrenAuthority(authority *models.Roles) (err error) {
	err = role_repository.GetChildrenList(authority.ID, authority)
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
//查询条件
func (r *RoleListService) GetWheres() map[string]interface{} {
	wheres := make(map[string]interface{})
	wheres["delete_status"] = 0
	if r.AuthorityName != "" {
		wheres["authority_name"] = r.AuthorityName
	} else {
		wheres["parent_id"] = 0
	}
	return wheres

}
