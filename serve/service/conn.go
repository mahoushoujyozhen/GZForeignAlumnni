package service

import (
	"backend/logging"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

var MSG_DATABASE_CONNECT = "the database connect success"
var ERR_DATABASE_CONNECT = errors.New("the database connect fail")

var MSG_SNOWFLAKE_CREATE = "the snowflake create success"
var ERR_SNOWFLAKE_CREATE = errors.New("the snowflake create fail")

var (
	connection *pgxpool.Pool
)

func init() {
	var errors error
	// 连接数据库
	config, err := pgxpool.ParseConfig("postgres://cstuser:cst4Ever@172.22.72.229:16900/kdb")
	if err != nil {
		logging.Error(err)
	}
	connection, errors = pgxpool.ConnectConfig(context.Background(), config)
	//处理连接失败的情况
	if errors != nil {
		logging.Error(ERR_DATABASE_CONNECT)
	} else {
		logging.Error(MSG_DATABASE_CONNECT)
	}
}

func GetDatabaseConnection() *pgxpool.Pool {
	return connection
}
