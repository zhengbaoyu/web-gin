package check_form

//公共分页
type PageInfo struct {
	Page     int `json:"page" form:"page" valid:"Required"`
	PageSize int `json:"pageSize" form:"pageSize" valid:"Required"`
}
