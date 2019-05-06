package e

var MsgFlags = map[int]string{
	SUCCESS: "请求成功",
	ERROR: "请求失败",
	INVALID_PARAMS: "参数有误",
	ERROR_EXIST_TAG: "标签已存在",
	ERROR_NOT_EXIST_TAG: "标签不存在",
	ERROR_NOT_EXIST_ARTICLE: "文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token 校验失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 校验超时",
	ERROR_AUTH_TOKEN: "Token 生成失败",
	ERROR_AUTH: "Token 有误",
}

func GetMsg(code int)  string{
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}