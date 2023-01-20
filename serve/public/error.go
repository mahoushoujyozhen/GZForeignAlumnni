package public

// ERROR_ 可参照的命名规范
//ERROR_EXIST_TAG = 10001
//ERROR_NOT_EXIST_TAG = 10002
//ERROR_NOT_EXIST_ARTICLE = 10003
//ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
//ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
//ERROR_AUTH_TOKEN = 20003
//ERROR_AUTH = 20004

const (
	//这几个我随便写的
	ERROR_NOT_RECEIVER   = 111
	ERROR_REQUEST        = 1111
	ERROR_READALL        = 111111
	ERR_SNOWFLAKE_CREATE = 1112
	MSG_SNOWFLAKE_CREATE = 1212
	//******

	SUCCESS                   = 200
	ERROR                     = 500
	INVALID_PARAMS            = 400
	ERROR_NOT_EXIST_USER      = 3000
	ERROR_NOT_EXIST_USER_ID   = 3001
	ERROR_NOT_EXIST_OPINION   = 3002
	ERROR_NOT_EXIST_STATUS    = 3003
	ERROR_NOT_EXIST_USER_INFO = 3004
	ERROR_ALREADY_EXISTS      = 4000
	ERR_SQL_CONNECT_MEMBER    = 4001
	ERR_SQL_CONNECT_EDU       = 4002
	ERR_SQL_CONNECT_WORK      = 4003

	ERROR_USER_ID = 3010
	ERROR_STATUS  = 3011

	ERROR_DATABASE_SEARCH_FAIL = 3020
	ERROR_DATABASE_UPDATE_FAIL = 3021

	MSG_SERVER_START = 504
	ERR_SERVER_START = 505

	CODE_HANDLE_SUCCESS = 2000
	CODE_HANDLE_FAILURE = 4000
)
