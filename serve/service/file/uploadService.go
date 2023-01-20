package file

import (
	"backend/logging"
	"backend/public"
	"backend/service"
	"backend/util"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func init() {
	err := os.Mkdir("/resources", 0666)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func GetUploadParams(request *http.Request) (string, string, multipart.File, *multipart.FileHeader, bool) {
	// 获取上传参数
	id := request.PostFormValue("id")
	md5Key := request.PostFormValue("md5Key")
	uploadFile, fileHeader, errors := request.FormFile("file")
	// 校验参数
	// "header.Size"获取文件大小且以字节(Byte)为单位
	if errors != nil || len(id) == 0 || len(md5Key) == 0 || uploadFile == nil || fileHeader == nil || fileHeader.Size > 10*2<<20 {
		return "", "", nil, nil, false
	} else {
		return id, md5Key, uploadFile, fileHeader, true
	}
}

func GetFilePath(uploadFileName string) (string, string, error) {
	var errors error
	// 获取文件后缀名
	extensionName := path.Ext(uploadFileName)
	// 生成随机文件名
	var fileName string
	fileName, errors = util.GenerateUniqueId(Snowflake)
	if errors != nil {
		return "", "", errors
	}
	return fmt.Sprintf("/resources/%v%v", fileName, extensionName), fmt.Sprintf("%v%v", fileName, extensionName), errors
	// return fmt.Sprintf("%v\\resources\\%v%v", rootPath, fileName, extensionName), fmt.Sprintf("%v%v", fileName, extensionName), errors

}

func UploadService(responseWriter http.ResponseWriter, request *http.Request) {
	var errors error
	// 校验参数
	var isValid bool
	var id, md5Key string
	var uploadFile multipart.File
	var fileHeader *multipart.FileHeader
	id, md5Key, uploadFile, fileHeader, isValid = GetUploadParams(request)
	if !isValid {
		service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "参数不合法", nil))
		return
	}
	// 获取数据库连接
	connection := service.GetDatabaseConnection()
	// 开启事务
	var transaction pgx.Tx
	transaction, errors = connection.Begin(context.Background())
	defer transaction.Rollback(context.Background())
	if errors != nil {
		logging.Error(errors.Error())
		service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "事务开启失败", nil))
		return
	}
	// 查询是否存在相同hash值的文件
	var fileName string
	fileName = ReadFileByMd5(md5Key)
	if len(fileName) == 0 {
		// 获取文件名
		var filePath string
		filePath, fileName, errors = GetFilePath(fileHeader.Filename)
		if errors != nil {
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "文件路径获取失败", nil))
			return
		}
		// 插入新文件的信息
		errors = InsertNewFile(md5Key, fileName)
		if errors != nil {
			logging.Error(errors.Error())
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "插入文件信息失败", nil))
			return
		}
		// 保存新文件
		var localFile *os.File
		localFile, errors = os.Create(filePath)
		if errors != nil {
			logging.Error(errors.Error())
			dir, _ := os.Getwd()
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "创建文件失败", dir))
			return
		}
		_, errors = io.Copy(localFile, uploadFile)
		defer localFile.Close()
		if errors != nil {
			logging.Error(errors.Error())
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "保存文件失败", nil))
			return
		}
	} else {
		// 存在相同文件,引用数加1
		errors = UpdateFileReferenceByMd5(md5Key)
		if errors != nil {
			logging.Error(errors.Error())
			service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "引用数更新失败", nil))
			return
		}
	}
	// 建立事件和文件的映射
	errors = InsertNewFileMapper(id, md5Key)
	if errors != nil {
		logging.Error(errors.Error())
		service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_FAILURE, "文件映射失败", nil))
		return
	}
	// 提交事务
	_ = transaction.Commit(context.Background())
	service.ResponseService(responseWriter, service.ResponseBeanFactory(public.CODE_HANDLE_SUCCESS, "success", fileName))
}
