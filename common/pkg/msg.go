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
	ERR_GET_MENU_LIST:              "获取权限功能列表失败",
	ERR_ADD_MENU_FAIL:              "添加权限功能失败",
	ERR_ADD_MENU_EXIST_NAME:        "存在重复name，请修改name",
	ERR_ADD_MENU_EXIST_TITLE:       "存在重复title，请修改title",
	ERR_VIEW_MENU_FAIL:             "获取详情失败",
	ERR_MENU_ID_NOT_EXIST:          "ID不存在",
	ERR_UPDATE_MENU_FAIL:           "编辑权限功能失败",
	ERR_DELETE_MENU_FAIL:           "删除失败",
	ERR_DELETE_MENU_EXIST:          "当前权限功能存在子权限,禁止删除",
	ERR_GET_ROLE_LIST:              "获取角色列表失败",
	ERR_ADD_ROLE_FAIL:              "添加角色失败",
	ERR_ROLE_NAME_EXIST:            "角色名称已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
