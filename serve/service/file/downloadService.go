package file

import (
	"backend/logging"
	"backend/public"
	"backend/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func DownloadService(responseWriter http.ResponseWriter, request *http.Request) {
	var errors error
	// 获取下载参数
	id := request.PostFormValue("id")
	fmt.Println("id", id)
	// 校验参数
	if len(id) == 0 {
		service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "fail", nil))
	} else {
		// 获取数据库连接
		connection := service.GetDatabaseConnection()
		// 开启事务
		var transaction pgx.Tx
		transaction, errors = connection.Begin(context.Background())
		if errors != nil {
			logging.Error(errors.Error())
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "fail", nil))
			return
		}
		// 查询事项所需要的所有文件
		var paths []string
		paths, errors = ReadFilePathById(id)
		if errors != nil {
			logging.Error(errors.Error())
			_ = transaction.Rollback(context.Background()) // 回滚
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "fail", nil))
		} else {
			_ = transaction.Commit(context.Background()) // 提交事务
			fmt.Println("success", paths)
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_SUCCESS, "success", paths))
		}
	}
}
