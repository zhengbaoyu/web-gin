package checkform

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-gin/common/pkg"
)

// 表单验证
func BindAndValidForm(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, pkg.INVALID_PARAMS
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, pkg.ERROR
	}

	if !check {
		return http.StatusBadRequest, pkg.INVALID_PARAMS
	}
	return http.StatusOK, pkg.SUCCESS
}
