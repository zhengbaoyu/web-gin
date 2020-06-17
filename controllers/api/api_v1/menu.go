package api_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/menu/addMenu [post]
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
	menuService := menu_service.Menu{
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
