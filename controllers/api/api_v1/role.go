package api_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-gin/common/pkg"
	"web-gin/common/pkg/checkform"
	"web-gin/controllers/api/api_v1/check_form"
	"web-gin/services/role_service"
)

// @Summary 角色列表
// @Produce  json
// @Param data body check_form.PageInfo true "角色搜索参数"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/role/list [get]
func GetRoleList(c *gin.Context) {
	var (
		response = pkg.Gin{C: c}
		form     check_form.RoleList
	)
	httpCode, Code := checkform.BindAndValidForm(c, &form)
	//参数校验
	if Code != pkg.SUCCESS {
		response.ResponseErr(httpCode, Code, nil)
		return
	}
	roleService := role_service.RoleListService{Page: form.Page, PageSize: form.PageSize, AuthorityName: form.AuthorityName}
	err, list, total := roleService.GetRoleList()
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_GET_ROLE_LIST, nil)
		return
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, gin.H{
		"list":     list,
		"total":    total,
		"page":     form.Page,
		"pageSize": form.PageSize,
	})
}

// @Summary 新增角色
// @Produce  json
// @Param data body check_form.AddMenu true "新增参数"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/role/add [post]
func AddRole(c *gin.Context) {

}

// @Summary 查看角色
// @Produce  json
// @Param id query string true "角色ID"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/role/view [get]
func ViewRole(c *gin.Context) {

}

// @Summary 修改角色
// @Produce  json
// @Param id query string true "角色ID"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/role/update [put]
func UpdateRole(c *gin.Context) {

}

// @Summary 删除角色
// @Produce  json
// @Param id query string true "角色ID"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/role/delete [delete]
func DeleteRole(c *gin.Context) {

}
