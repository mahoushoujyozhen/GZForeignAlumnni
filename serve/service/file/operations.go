package file

import (
	"backend/service"
	"backend/util"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func ReadFileByMd5(md5Key string) string {
	// SQL语句
	readSQL := `
		SELECT path
		FROM t_file
		WHERE hash = $1;
	`
	// 执行SQL语句
	results := dbConn.QueryRow(context.Background(), readSQL, md5Key)
	// 返回执行结果
	var path string
	_ = results.Scan(&path)
	return path
}

func UpdateFileReferenceByMd5(md5Key string) error {
	var errors error
	// SQL语句
	updateSQL := `
		UPDATE t_file
		SET reference = t_file.reference + 1
		WHERE hash = $1;
	`
	// 执行SQL语句
	_, errors = service.GetDatabaseConnection().Exec(context.Background(), updateSQL, md5Key)
	return errors
}

func InsertNewFile(md5Key string, filePath string) error {
	var errors error
	// SQL语句
	insertSQL := `
		INSERT INTO t_file
		(hash, path, reference)
		VALUES
		($1, $2, 1);
	`
	// 执行SQL语句
	_, errors = service.GetDatabaseConnection().Exec(context.Background(), insertSQL, md5Key, filePath)
	fmt.Println("文件上传err", errors)
	// 返回执行结果
	return errors
}

func InsertNewFileMapper(id string, md5Key string) error {
	var errors error
	// SQL语句
	insertSQL := `
		INSERT INTO t_file_mapper
		(id, hash)
		VALUES
		($1, $2);
	`
	// 执行SQL语句
	_, errors = service.GetDatabaseConnection().Exec(context.Background(), insertSQL, id, md5Key)
	// 返回执行结果
	return errors
}

func InsertNewMessage(uid string, title string, time int64, content string) error {
	var errors error
	// 生成唯一Id
	var mid string
	mid, errors = util.GenerateUniqueId(Snowflake)
	// SQL语句
	insertSQL := `
		INSERT INTO t_message
		(mid, uid, title, time, content)
		VALUES
		($1, $2, $3, $4, $5);
	`
	// 执行SQL语句
	_, errors = service.GetDatabaseConnection().Exec(context.Background(), insertSQL, mid, uid, title, time, content)
	return errors
}

func ReadMessageByUserId(uid string) ([]service.NotificationBean, error) {
	var errors error
	// SQL语句
	readSQL := `
		SELECT mid, title, time, content
		FROM t_message
		WHERE uid = $1;
	`
	// 执行SQL语句
	var results pgx.Rows
	results, errors = service.GetDatabaseConnection().Query(context.Background(), readSQL, uid)
	if errors != nil {
		return nil, errors
	}
	// 扫描执行结果
	var time int64
	var mid, title, content string
	var notifications []service.NotificationBean
	for results.Next() {
		errors = results.Scan(&mid, &title, &time, &content)
		if errors != nil {
			return nil, errors
		}
		notifications = append(notifications, service.NotificationBean{Id: mid, Title: title, Time: time, Content: content})
	}
	return notifications, errors
}

func DeleteMessageByUserId(uid string) error {
	var errors error
	// SQL语句
	deleteSQL := `
		DELETE FROM t_message
		WHERE uid = $1;
	`
	// 执行SQL语句
	_, errors = service.GetDatabaseConnection().Exec(context.Background(), deleteSQL, uid)
	return errors
}

func ReadFilePathById(id string) ([]string, error) {
	var errors error
	// SQL语句
	readSQL := `
		SELECT path
		FROM t_file
		WHERE 
		hash IN
		(
			SELECT hash
			FROM t_file_mapper
			WHERE id = $1
		);
	`
	// 执行SQL语句
	var results pgx.Rows
	results, errors = service.GetDatabaseConnection().Query(context.Background(), readSQL, id)
	if errors != nil {
		return nil, errors
	}
	// 扫描执行结果
	var path string
	var paths []string
	for results.Next() {
		errors = results.Scan(&path)
		if errors != nil {
			return nil, errors
		}
		paths = append(paths, path)
	}
	return paths, errors
}

func DeleteFileMapperById(id string) error {

	var errors error
	// SQL语句
	deleteSQL := `
		DELETE FROM t_file_mapper
		WHERE id = $1;
	`
	// 执行SQL语句
	_, errors = service.GetDatabaseConnection().Exec(context.Background(), deleteSQL, id)
	return errors
}

func DeleteNullReferenceFile() ([]string, error) {
	var errors error
	// SQL语句
	readSQL := `
		DELETE FROM t_file
		WHERE reference = 0
		RETURNING path;
	`
	// 执行SQL语句
	var results pgx.Rows
	results, errors = service.GetDatabaseConnection().Query(context.Background(), readSQL)
	// 扫描执行结果
	var fileName string
	var fileNames []string
	for results.Next() {
		errors = results.Scan(&fileName)
		if errors != nil {
			return nil, errors
		}
		fileNames = append(fileNames, fileName)
	}
	return fileNames, errors
}
