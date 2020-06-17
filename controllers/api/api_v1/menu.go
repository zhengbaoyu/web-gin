package api_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-gin/common/pkg"
	"web-gin/common/pkg/checkform"
	"web-gin/controllers/api/api_v1/check_form"
	"web-gin/services/menu_service"
)

// @Summary 获取权限功能列表
// @Produce  json
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/menu/list [get]
func GetMenuList(c *gin.Context) {
	response := pkg.Gin{C: c}
	err, menuList := menu_service.GetInfoList()
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_GET_MENU_LIST, err)
		return
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, gin.H{
		"list": menuList,
	})
}

// @Summary 新增权限功能
// @Produce  json
// @Param data body check_form.AddMenu true "新增权限功能"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/menu/add [post]
func AddMenu(c *gin.Context) {
	var (
		response = pkg.Gin{C: c}
		form     check_form.AddMenu
	)
	httpCode, Code := checkform.BindAndValidForm(c, &form)
	//参数校验
	if Code != pkg.SUCCESS {
		response.ResponseErr(httpCode, Code, nil)
		return
	}
	menuService := menu_service.AddMenuService{
		ParentId:  form.ParentId,
		Path:      form.Path,
		Name:      form.Name,
		Hidden:    form.Hidden,
		Component: form.Component,
		Sort:      form.Sort,
		Title:     form.Title,
		Icon:      form.Icon,
	}
	errCode := menuService.AddMenu()
	if errCode != pkg.SUCCESS {
		response.ResponseErr(http.StatusInternalServerError, errCode, nil)
		return
	}
	//返回数据
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, nil)
}

// @Summary 获取权限功能详情
// @Produce  json
// @Param id query string true "权限功能ID"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/menu/view/:id [get]
func ViewMenu(c *gin.Context) {
	response := pkg.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_MENU_ID_NOT_EXIST, nil)
		return
	}
	menudata, err := menu_service.ViewMenu(uint(id))
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_VIEW_MENU_FAIL, err)
		return
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, gin.H{
		"view": menudata,
	})
}

// @Summary 修改权限菜单
// @Produce  json
// @Param id query string true "权限功能ID"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/menu/update/:id [put]
func UpdateMenu(c *gin.Context) {
	var (
		response = pkg.Gin{C: c}
		form     check_form.UpdateMenu
		id, _    = strconv.Atoi(c.Param("id"))
	)
	httpCode, Code := checkform.BindAndValidForm(c, &form)
	//参数校验
	if Code != pkg.SUCCESS {
		response.ResponseErr(httpCode, Code, nil)
		return
	}
	updateService := menu_service.UpdateMenuService{
		Id:        uint(id),
		ParentId:  form.ParentId,
		Path:      form.Path,
		Name:      form.Name,
		Hidden:    form.Hidden,
		Component: form.Component,
		Sort:      form.Sort,
		Title:     form.Title,
		Icon:      form.Icon,
	}
	errcode := updateService.UpdateMenu()
	if errcode != pkg.SUCCESS {
		response.ResponseErr(http.StatusInternalServerError, errcode, nil)
		return
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, nil)
}

// @Summary 删除权限菜单
// @Produce  json
// @Param id query string true "权限功能ID"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/menu/delete/:id [delete]
func DeleteMenu(c *gin.Context) {
	response := pkg.Gin{C: c}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, nil)
}
