package pkg

var MsgFlags = map[int]string{
	SUCCESS:                        "OK",
	ERROR:                          "ERROR",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_PARSE:   "Token解析失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token超时",
	ERR_PASSWORD_NOT_EQUAL:         "两次输入的密码不一致",
	ERR_REGISTER_FAIL:              "注册失败",
	ERR_EXIST_PHONE:                "手机号已存在",
	ERR_NOT_EXIST_PHONE:            "该手机号未注册",
	ERR_PASSWORD_FAIL:              "密码错误",
	ERR_GET_TOKEN_FAIL:             "获取token失败",
	ERR_GET_USER:                   "获取用户信息失败",
	ERR_GET_USER_NULL:              "用户不存在",
	ERR_DOWNLOAD_FAIL:              "下载失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
