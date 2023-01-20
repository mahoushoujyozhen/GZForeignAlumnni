package file

import (
	"backend/logging"
	"backend/public"
	"backend/service"
	"backend/util"
	"net/http"
)

func GenerateIdService(responseWriter http.ResponseWriter, request *http.Request) {

	// 生成文件Id
	fileId, errors := util.GenerateUniqueId(Snowflake)
	if errors != nil {
		logging.Error(errors.Error())
		service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "fail", nil))
	} else {
		service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_SUCCESS, "success", fileId))
	}
}
