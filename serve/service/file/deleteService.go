package file

import (
	"backend/logging"
	e2 "backend/public"
	"backend/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
	"time"
)

func init() {
	go DeleteNullReferenceFiles()
}

func DeleteService(responseWriter http.ResponseWriter, request *http.Request) {
	var errors error
	// 获取请求参数
	id := request.PostFormValue("id")
	// 校验参数
	if len(id) == 0 {
		service.ResponseService(responseWriter, service.ResponseBeanFactory(e2.CODE_HANDLE_FAILURE, e2.GetMsg(e2.INVALID_PARAMS), nil))
	} else {
		// 获取数据库连接
		connection := service.GetDatabaseConnection()
		// 开启事务
		var transaction pgx.Tx
		transaction, errors = connection.Begin(context.Background())
		if errors != nil {
			logging.Error(errors.Error())
			service.ResponseService(responseWriter, service.ResponseBeanFactory(e2.CODE_HANDLE_FAILURE, "fail", nil))
			return
		}
		// 删除事件和文件的映射
		errors = DeleteFileMapperById(id)
		if errors != nil {
			logging.Error(errors.Error())
			_ = transaction.Rollback(context.Background()) // 回滚
			service.ResponseService(responseWriter, service.ResponseBeanFactory(e2.CODE_HANDLE_FAILURE, "fail", nil))
			return
		}
		_ = transaction.Commit(context.Background()) // 提交事务
		service.ResponseService(responseWriter, service.ResponseBeanFactory(e2.CODE_HANDLE_SUCCESS, "success", nil))
	}
}

// DeleteNullReferenceFiles 定时检查引用数为0的文件
func DeleteNullReferenceFiles() {
	timer := time.NewTimer(time.Minute * 1)
	for {
		select {
		case <-timer.C:
			var errors error
			// 获取引用数为0的文件名
			var fileNames []string
			fileNames, errors = DeleteNullReferenceFile()
			if errors != nil {
				logging.Error(errors.Error())
				return
			}
			// 删除对应的文件
			for i := 0; i < len(fileNames); i++ {
				errors = os.Remove(fmt.Sprintf("/resources/%v", fileNames[i]))
				if errors != nil {
					logging.Error(errors.Error())
					return
				}
			}
			// 重置定时器
			timer.Reset(time.Minute * 1)
		}
	}
}
