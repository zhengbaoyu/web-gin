package check_form

type RoleList struct {
	PageInfo
	AuthorityName string `form:"authorityName"`
}
