package check_form

type RoleList struct {
	PageInfo
	AuthorityName string `form:"authorityName"`
}

type AddRoleForm struct {
	ParentId      int `form:"parentId"`
	AuthorityName string `form:"authorityName"`
}
