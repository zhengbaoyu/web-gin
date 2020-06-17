package api_v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"web-gin/common/pkg"
	"web-gin/common/pkg/checkform"
	"web-gin/common/pkg/jwt"
	"web-gin/controllers/api/api_v1/check_form"
	"web-gin/models/return_data"
	"web-gin/services/user_service"
)

func Register(c *gin.Context) {
	var (
		response = pkg.Gin{C: c}
		form     check_form.RegisterForm
	)
	httpCode, Code := checkform.BindAndValidForm(c, &form)
	if Code != pkg.SUCCESS {
		response.ResponseErr(httpCode, Code, nil)
		return
	}
	if form.Password != form.RePassword {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_PASSWORD_NOT_EQUAL, nil)
		return
	}
	//密码加密
	password, _ := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	userService := user_service.User{
		Phone:    form.Phone,
		Password: string(password),
	}
	//校验手机号是否存在
	userData, err := userService.CheckPhone()
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERROR, err)
		return
	}
	if userData.Id > 0 {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_EXIST_PHONE, nil)
		return
	}
	//创建
	err = userService.AddUser()
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_REGISTER_FAIL, nil)
		return
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, "注册成功")
}
// @Summary 用户登录
// @Produce  json
// @Param data body check_form.LoginForm true "用户登录接口"
// @Success 200 {string} string "{"status":true,"returnCode":200,"msg":"OK",data":{}}"
// @Router /v1/user/login [post]
func Login(c *gin.Context) {
	var (
		response = pkg.Gin{C: c}
		form     check_form.LoginForm
	)
	httpCode, Code := checkform.BindAndValidForm(c, &form)
	if Code != pkg.SUCCESS {
		response.ResponseErr(httpCode, Code, nil)
		return
	}
	userService := user_service.User{
		Phone:    form.Phone,
		Password: form.Password,
	}
	userData, err := userService.CheckPhone()
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERROR, err)
		return
	}
	if userData.Id == 0 {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_NOT_EXIST_PHONE, nil)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(form.Password)); err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_PASSWORD_FAIL, nil)
		return
	}
	returnUser := return_data.ReturnLoginUser{
		Id:       userData.Id,
		RoleId:   userData.RoleId,
		Name:     userData.Name,
		Nickname: userData.Nickname,
		Phone:    userData.Phone,
		Sex:      userData.Sex,
		Age:      userData.Age,
		Avatar:   userData.Avatar,
	}
	//取token并加密数据
	token, err := jwt.ReleaseToken(userData.Id)
	if err != nil {
		response.ResponseErr(http.StatusInternalServerError, pkg.ERR_GET_TOKEN_FAIL, nil)
		return
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, gin.H{
		"token": token,
		"list":  returnUser,
	})
}
func GetUser(c *gin.Context) {
	response := pkg.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	userService := user_service.User{
		Id: id,
	}
	userData, err := userService.GetUser()
	if err != nil {
		response.ResponseOk(http.StatusOK, pkg.ERR_GET_USER, err)
		return
	}
	if userData.Id == 0 {
		response.ResponseOk(http.StatusOK, pkg.ERR_GET_USER_NULL, err)
		return
	}
	returnUser := return_data.ReturnGetUser{
		Id:        userData.Id,
		RoleId:    userData.RoleId,
		Name:      userData.Name,
		Nickname:  userData.Nickname,
		Phone:     userData.Phone,
		Sex:       userData.Sex,
		Age:       userData.Age,
		Avatar:    userData.Avatar,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
	response.ResponseOk(http.StatusOK, pkg.SUCCESS, returnUser)
}
