package public

//  命名规范
//	ERROR_EXIST_TAG : "已存在该标签名称",
//	ERROR_NOT_EXIST_TAG : "该标签不存在",
//	ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
//	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
//	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
//	ERROR_AUTH_TOKEN : "Token生成失败",
//	ERROR_AUTH : "Token错误",

var MsgFlags = map[int]string{

	//消息模块
	ERROR_NOT_RECEIVER: "查无此收件人",
	ERROR_REQUEST:      "invalid request",
	ERROR_READALL:      "读取浏览器请求信息错误",

	SUCCESS:                "ok",
	ERROR:                  "fail",
	INVALID_PARAMS:         "请求参数错误",
	ERROR_NOT_EXIST_USER:   "查找用户错误",
	MSG_SERVER_START:       "服务启动成功",
	ERR_SERVER_START:       "服务启动失败",
	ERR_SQL_CONNECT_MEMBER: "Member数据库更新错误",
	ERR_SQL_CONNECT_EDU:    "edu数据库更新错误",
	ERR_SQL_CONNECT_WORK:   "work数据库更新错误",
	ERROR_ALREADY_EXISTS:   "用户已存在",

	ERROR_USER_ID: "用户ID应该为数字",
	ERROR_STATUS:  "审核状态应该为数字",

	ERROR_DATABASE_SEARCH_FAIL: "数据库查询错误",
	ERROR_DATABASE_UPDATE_FAIL: "数据库更新错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
