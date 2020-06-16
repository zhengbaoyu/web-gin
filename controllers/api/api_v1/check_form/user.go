package check_form

//注册参数校验
type RegisterForm struct {
	Phone      string `form:"phone" valid:"Required"`
	Password   string `form:"password" valid:"Required"`
	RePassword string `form:"repassword" valid:"Required"`
}

//登录参数校验
type LoginForm struct {
	Phone      string `form:"phone" valid:"Required"`
	Password   string `form:"password" valid:"Required"`
}
